package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
	"qqbot/handle/dto"
	"qqbot/handle/reward"
	"qqbot/handle/tb"
	"qqbot/handle/util"
	"qqbot/pkg/db"
)

func ReceivePet(uid string) (res *tb.Pet, err error) {
	// 判断用户是否有宠物
	var count int64
	if e := db.DB.Model(&tb.UserPet{}).Where(&tb.UserPet{UID: uid}).Count(&count).Error; e != nil {
		return nil, e
	}
	if count > 0 {
		return nil, errors.New("🐱 不可重复领取初始宠物")
	}

	// 查询初始宠物信息
	var pet tb.Pet
	if e := db.DB.Model(&pet).Preload("Feature").First(&pet, 137).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	if pet.ID == 0 {
		return nil, errors.New("❌ can not find id=1 pet")
	}

	// 发放宠物
	if e := reward.NewUserReward(uid).AddPets(&tb.UserPet{
		UID:   uid,
		PetId: 137,
		ChaId: util.RandCharacterId(),
	}).Save(); e != nil {
		return nil, e
	}

	res = &pet
	return
}

func PetBag(uid string, pageNo int) (res *dto.PetBag, err error) {
	// 查询用户宠物总数
	var count int64
	if e := db.DB.Model(&tb.UserPet{}).Where(tb.UserPet{UID: uid}).Count(&count).Error; e != nil {
		return nil, e
	}

	// 计算总页数
	pageSize := 10
	pageTotal := int(math.Ceil(float64(count) / float64(pageSize)))

	// 分页查询用户宠物
	var pets []*tb.UserPet
	if e := db.DB.Model(&tb.UserPet{}).Preload("Pet").Preload("Cha").Where(&tb.UserPet{UID: uid}).Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("lv desc, serial asc").Find(&pets).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	if len(pets) == 0 {
		return nil, errors.New("🐱 背包空空如也")
	}

	res = &dto.PetBag{
		PageNo:    pageNo,
		PageTotal: pageTotal,
		Pets:      pets,
	}
	return
}

func QueryUserPetBySerial(uid string, serial uint) (res *tb.UserPet, err error) {
	pet, err := GetUserPetBySerial(uid, serial)
	if err != nil {
		return nil, err
	}
	if pet.ID == "" {
		return nil, fmt.Errorf("🐱 没有找到ID为%v的宠物", serial)
	}
	return pet, nil
}

func QueryUserPetByName(uid string, name string) (res *tb.UserPet, err error) {
	pet, err := GetUserPetByName(uid, name)
	if err != nil {
		return nil, err
	}
	if pet.ID == "" {
		return nil, fmt.Errorf("🐱 没有找到名称为%v的宠物", name)
	}
	return pet, nil
}

func QueryPetById(id uint) (res *tb.Pet, err error) {
	pet, err := GetPetById(id)
	if err != nil {
		return nil, err
	}
	if pet.ID == 0 {
		return nil, fmt.Errorf("📓 没有找到ID为%v的宠物", id)
	}
	return pet, nil
}

func QueryPetByName(name string) (res *tb.Pet, err error) {
	pet, err := GetPetByName(name)
	if err != nil {
		return nil, err
	}
	if pet.ID == 0 {
		return nil, fmt.Errorf("📓 没有找到名称为%v的宠物", name)
	}
	return pet, nil
}

func GetUserPetById(id string) (res *tb.UserPet, err error) {
	// 根据序号查询用户宠物
	var pet tb.UserPet
	if e := db.DB.Model(&pet).Preload("Pet.Feature").Preload("Pet.EvoPet").Preload("Cha").First(&pet, id).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	return &pet, nil
}

func GetUserPetBySerial(uid string, serial uint) (res *tb.UserPet, err error) {
	// 根据序号查询用户宠物
	var pet tb.UserPet
	if e := db.DB.Model(&pet).Preload("Pet.Feature").Preload("Pet.EvoPet").Preload("Cha").Where(&tb.UserPet{UID: uid, Serial: serial}).First(&pet).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	return &pet, nil
}

func GetUserPetByName(uid string, name string) (res *tb.UserPet, err error) {
	// 根据名称查询用户宠物
	var pet tb.UserPet
	if e := db.DB.Model(&pet).Joins("Pet", "Pet.Feature", "Pet.EvoPet").Preload("Cha").Where("user_pets.uid = ? AND pet.name = ?", uid, name).First(&pet).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	return &pet, nil
}

func GetPetById(id uint) (res *tb.Pet, err error) {
	// 根据ID查询宠物
	var pet tb.Pet
	if e := db.DB.Model(&pet).Preload("Feature").Preload("EvoPet").First(&pet, id).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	return &pet, nil
}

func GetPetByName(name string) (res *tb.Pet, err error) {
	// 根据名称查询宠物
	var pet tb.Pet
	if e := db.DB.Model(&pet).Preload("Feature").Preload("EvoPet").Where(&tb.Pet{Name: name}).First(&pet).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	return &pet, nil
}
