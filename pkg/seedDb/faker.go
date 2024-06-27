package seedDb

import "github.com/jaswdr/faker/v2"

// TODO: Had to put this here otherwise it will not give it a random seed properly,
// need to find a better place for it
var fake = faker.New()
