package buffalo

type Repository interface{
	FindAll()([]Buffalo, error)
	Create()
}