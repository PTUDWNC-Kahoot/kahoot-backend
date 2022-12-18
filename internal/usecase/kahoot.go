package usecase


type kahootUsecase struct {
	repo KahootRepo
}

func NewKahootUsecase(repo KahootRepo) KahootUsecase {
	return &kahootUsecase{
		repo: repo,
	}

}
