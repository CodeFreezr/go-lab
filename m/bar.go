package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Bar() {

	emojitree := `
┌────────────────┐
│ emoji v5 (all) │
├────────────────┘
│              
├────────────────────────────┬────────────────────────────┬─────────────────────────────── 4. TravelPlaces 🚀        
├── 1. SmileyPeoples 🙂       └── 2. AnimalsNature 🐸       └── 3. FoodDrink ☕               ├── hotel               
│   ├── body                     ├── animal-amphibian         ├── dishware                 ├── place-building      
│   ├── cat-face                 ├── animal-bird              ├── drink                    ├── place-geographic    
│   ├── clothing                 ├── animal-bug               ├── food-asian               ├── place-map           
│   ├── emotion                  ├── animal-mammal            ├── food-fruit               ├── place-other         
│   ├── face-fantasy             ├── animal-marine            ├── food-prepared            ├── place-religious     
│   ├── face-negative            ├── animal-reptile           ├── food-sweet               ├── sky-weather         
│   ├── face-neutral             ├── plant-flower             └── food-vegetable           ├── time                
│   ├── face-positive            └── plant-other                                           ├── transport-air       
│   ├── face-role                                                                          ├── transport-ground    
│   ├── face-sick                                                                          └── transport-water     
│   ├── family                                                                                                     
│   ├── monkey-face                                                                                                
│   ├── person                                                                                                     
│   ├── person-activity                                                                                            
│   ├── person-fantasy                                                                                             
│   ├── person-gesture                                                                                             
│   ├── person-role                                                                                                
│   └── person-sport
│
├────────────────────────────┬────────────────────────────┬─────────────────────────────── 8. Flags 🏳️‍                
└── 5. Activities 🏆          └── 6. Objects 🔑             └── 7. Symbols ♻                 ├── country-flag        
    ├── award-medal              ├── book-paper               ├── alphanum                 ├── flag                
    ├── event                    ├── computer                 ├── arrow                    └── subdivision-flag    
    ├── game                     ├── light-video              ├── av-symbol                
    └── sport                    ├── lock                     ├── geometric            
                                 ├── mail                     ├── keycap               
                                 ├── money                    ├── other-symbol         
                                 ├── music                    ├── religion             
                                 ├── musical-instrument       ├── transport-sign       
                                 ├── office                   ├── warning              
                                 ├── other-object             └── zodiac               
                                 ├── phone                 
                                 ├── sound                 
                                 ├── tool                  
                                 └── writing`

	fmt.Println(emojitree)
}
