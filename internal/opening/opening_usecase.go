package opening

import "github.com/google/uuid"

type OpeningUseCase struct {
	storage *OpeningStorage
}

func newOpeningUseCase(storage *OpeningStorage) *OpeningUseCase {
	return &OpeningUseCase{
		storage: storage,
	}
}

func (useCase *OpeningUseCase) createOpening(request *CreateOpeningRequest) (*Opening, error) {
	schemaObj := request.TransformRequestToSchema()

	err := useCase.storage.createOpening(schemaObj)
	if err != nil {
		return nil, err
	}

	return schemaObj, nil
}

func (useCase *OpeningUseCase) deleteOpening(id string) error {
	idFormated, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return useCase.storage.deleteOpening(&idFormated)
}

func (useCase *OpeningUseCase) listOpenings() ([]*OpeningResponseDTO, error) {
	return useCase.storage.listOpenings()
}

func (useCase *OpeningUseCase) showOpening(id string) (*OpeningResponseDTO, error) {
	idFormated, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return useCase.storage.showOpening(idFormated)
}

func (useCase *OpeningUseCase) updateOpening(id string, request *CreateOpeningRequest) (*Opening, error) {
	var err error

	idFormated, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	schemaObj := request.TransformRequestToSchema()

	err = useCase.storage.updateOpening(&idFormated, schemaObj)
	if err != nil {
		return nil, err
	}

	return schemaObj, nil
}
