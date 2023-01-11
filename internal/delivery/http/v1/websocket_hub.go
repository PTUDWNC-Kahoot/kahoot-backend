package v1

import (
	"encoding/json"
	"log"
	"sort"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Broadcast struct {
	roomID   uint32
	userID   uint32
	Username string
	data     []byte
}

const questionTimer = time.Second * 15 // 15s
const pointsPosible = 1000.0

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Broadcast

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister    chan *Client
	presentations map[uint32]map[string]interface{}
}

func newHub() *Hub {
	return &Hub{
		broadcast:     make(chan Broadcast),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		presentations: make(map[uint32]map[string]interface{}),
	}
}

type Message struct {
	Action  string      `json:"action"`
	Code    string      `json:"code" `
	Payload interface{} `json:"payload"`
}

type Ranking struct {
	UserID   uint32 `json:"userId"`
	Username string `json:"username"`
	Score    int32  `json:"score"`
}

type SubmitAnswerPayload struct {
	AnswerID  uint32 `json:"answerId"`
	IsCorrect bool   `json:"isCorrect"`
}

type CurrentSlide struct {
	Deadline time.Time `json:"deadline"`
	SlideID  uint32    `json:"slideId"`
	Index    int32     `json:"index"`
	Anwers   map[uint32]int32
}

type CurrentSlidePayload struct {
	Index   int32  `json:"index"`
	SlideID uint32 `json:"slideId"`
}

type ResultPayload struct {
	Ranking []*Ranking       `json:"ranking"`
	Answers map[uint32]int32 `json:"answers"`
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case b := <-h.broadcast:
			m := Message{}
			err := json.Unmarshal(b.data, &m)
			if err != nil {
				log.Println("failed to unmarshal data", err)
				continue
			}

			if h.presentations[b.roomID] == nil {
				h.presentations[b.roomID] = map[string]interface{}{}
			}

			if m.Action == "goto_slide" {
				var payload CurrentSlidePayload
				err := mapstructure.Decode(m.Payload, &payload)
				if err != nil {
					log.Println("failed to decode payload")
				}

				h.presentations[b.roomID]["current_slide"] = &CurrentSlide{
					Deadline: time.Now().Add(questionTimer),
					SlideID:  payload.SlideID,
					Index:    payload.Index,
				}

				// setTimeout for response result
				go h.broadcastResult(b.roomID)
			}

			if m.Action == "submit_answer" {
				ranking, ok := h.presentations[b.roomID]["ranking"].(map[uint32]*Ranking)
				if !ok {
					ranking = make(map[uint32]*Ranking)
				}

				currentSlide, ok := h.presentations[b.roomID]["current_slide"].(*CurrentSlide)
				if !ok {
					log.Println("failed to get current slide", err)
					continue
				}

				var payload SubmitAnswerPayload
				err := mapstructure.Decode(m.Payload, &payload)
				if err != nil {
					log.Println("failed to get submit answer payload", err)
					continue
				}

				if currentSlide.Anwers == nil {
					currentSlide.Anwers = make(map[uint32]int32)
				}

				currentSlide.Anwers[payload.AnswerID] += 1

				// handle score
				untilTime := time.Until(currentSlide.Deadline).Seconds()
				if untilTime < 0 {
					untilTime = 0
				}

				responseTime := questionTimer.Seconds() - untilTime

				log.Printf("response time: %f", responseTime)

				subtractBy := 1.0
				if !payload.IsCorrect {
					subtractBy = 0.7
				}

				score := (subtractBy - ((responseTime / questionTimer.Seconds()) / 2)) * pointsPosible
				log.Printf("score: %f", score)

				if ranking[b.userID] == nil {
					ranking[b.userID] = &Ranking{
						Score:    0,
						UserID:   b.userID,
						Username: b.Username,
					}
				}

				ranking[b.userID].Score += int32(score)

				h.presentations[b.roomID]["ranking"] = ranking
			}

			for client := range h.clients {
				if client.roomID == b.roomID {
					select {
					case client.send <- b.data:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}

			}
		}
	}
}

func (h *Hub) broadcastResult(roomID uint32) {
	time.Sleep(questionTimer + time.Second)
	ranking, ok := h.presentations[roomID]["ranking"].(map[uint32]*Ranking)
	if !ok {
		log.Println("failed to cast ranking data")
		return
	}
	result := ResultPayload{
		Ranking: make([]*Ranking, 0, len(ranking)),
		Answers: map[uint32]int32{},
	}

	for _, rank := range ranking {
		result.Ranking = append(result.Ranking, rank)
	}

	sort.Slice(result.Ranking, func(i, j int) bool {
		return result.Ranking[i].Score > result.Ranking[j].Score
	})

	// get current answers stats
	currentSlide, ok := h.presentations[roomID]["current_slide"].(*CurrentSlide)
	if !ok {
		log.Println("failed to get current slide")
		return
	}

	result.Answers = currentSlide.Anwers

	data, err := json.Marshal(Message{
		Action:  "show_result",
		Payload: result,
	})
	if err != nil {
		log.Println("failed to marshal result", err)
		return
	}

	h.broadcast <- Broadcast{
		roomID: roomID,
		data:   data,
	}
}
