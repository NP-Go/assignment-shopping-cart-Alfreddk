package main

import (
	"fmt"
	"strconv"
	"strings"
)

var categories = []string{"Household", "Food", "Drinks"}

type itemInformation struct {
	categoryOfItem int
	quantityOfItem int
	unitCostOfItem float64
}

var item = map[string]itemInformation{ //preloaded items
	"Fork":   {0, 4, 3},
	"Plates": {0, 4, 3},
	"Cups":   {0, 5, 3},
	"Bread":  {1, 2, 2},
	"Cake":   {1, 3, 1},
	"Coke":   {2, 5, 2},
	"Sprite": {2, 5, 2},
}

var shoppingList = make([]map[string]itemInformation, 3) //can we have varying length on the inner slice for each elements?
var categoriesSaveFile = make([][]string, 3)             //can we have varying length on the inner slice for each elements?

var shoppingListMenu = []string{
	"\n",
	"Shopping List Application\n",
	"=========================\n",
	"1. View entire shopping list.\n",
	"2. Generate Shopping List Report.\n",
	"3. Add Items.\n",
	"4. Modify Items.\n",
	"5. Delete Item.\n",
	"6. Print Current Data.\n",
	"7. Add New Category Name.\n",
	"8. Modify category.\n",
	"9. Delete category.\n",
	"10. Storing of shopping lists.\n",
	"11. Retrieving of lists.\n",
	"Select your choice:",
}

var generateReportMenu = []string{
	"Generate Report\n",
	"1. Total Cost of each category.\n",
	"2. List of item by category.\n",
	"3. Main Menu.\n",
	"\n",
	"Choose your report:",
	"\n",
}

func stringSliceContainCheck(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func stringSliceIndexCheck(s []string, str string) (index int) {
	for k, v := range s {
		if v == str {
			index = k
		}
	}
	return index
}

func removeSliceIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	for {

		var shoppingListMenuSelection int
		var generateReportMenuSelection int

		fmt.Println(strings.Join(shoppingListMenu, ""))
		fmt.Scanln(&shoppingListMenuSelection)

		switch shoppingListMenuSelection {
		case 1:
			fmt.Println("Shopping List Contents:")
			for k, v := range item {
				cat := categories[v.categoryOfItem]
				qty := strconv.Itoa(v.quantityOfItem)
				uCost := strconv.FormatFloat(v.unitCostOfItem, 'f', -1, 64)

				fmt.Println("Category: " + cat + " - Item: " + k + " Quantity: " + qty + " Unit Cost: " + uCost)
			}

		case 2:
			for {
				fmt.Println(strings.Join(generateReportMenu, ""))
				fmt.Scanln(&generateReportMenuSelection)
				switch generateReportMenuSelection {
				case 1:
					fmt.Println("Total Cost of each Category.")
					for i := range categories {
						totalCostByCat := make([]float64, len(categories))
						for _, v := range item {
							if v.categoryOfItem == i {
								totalCostByCat[i] = totalCostByCat[i] + (float64(v.quantityOfItem) * v.unitCostOfItem)
							} else {
								continue
							}
						}
						fmt.Println(categories[i] + " cost : " + strconv.FormatFloat(totalCostByCat[i], 'f', -1, 64))
					}

				case 2:
					fmt.Println("List by Category.")
					for i := range categories {
						for k, v := range item {
							if v.categoryOfItem == i {
								cat := categories[v.categoryOfItem]
								qty := strconv.Itoa(v.quantityOfItem)
								uCost := strconv.FormatFloat(v.unitCostOfItem, 'f', -1, 64)
								fmt.Println("Category: " + cat + " - Item: " + k + " Quantity: " + qty + " Unit Cost: " + uCost)
							} else {
								continue
							}
						}
					}

				case 3:
					break

				default:
					fmt.Println("Invalid input.")
					continue
				}
				break
			}

		case 3:
			for {
				fmt.Println("Add item.")

				var itemName string
				var catName string
				var catIndex int
				var qty int
				var uCost float64

				fmt.Println("What is the name of your item?")
				fmt.Scanln(&itemName)

				_, okItemName := item[itemName]
				if okItemName == true {
					fmt.Println("Item already exist!")
					continue
				} else {
					fmt.Println("What category does it belong to?")
					fmt.Scanln(&catName)
				}

				if stringSliceContainCheck(categories, catName) == false {
					fmt.Println("This category does not exist!")
					continue
				} else {
					catIndex = stringSliceIndexCheck(categories, catName)
				}

				fmt.Println("How many units are there")
				fmt.Scanln(&qty)

				fmt.Println("How much does it cost per unit?")
				fmt.Scanln(&uCost)

				item[itemName] = itemInformation{catIndex, qty, uCost}

				break
			}

		case 4:
			for {
				fmt.Println("Modify Item.")
				var itemName string
				var catName string
				//var catIndex int
				var qty int
				var uCost float64

				fmt.Println("What item would you wish to modify")
				fmt.Scanln(&itemName)

				v, okItemName := item[itemName]
				if okItemName == false {
					fmt.Println("Item does not exist!")
					continue
				} else {
					catName = categories[v.categoryOfItem]
					//catIndex = v.categoryOfItem
					qty = v.quantityOfItem
					uCost = v.unitCostOfItem
					fmt.Println("Current item name is " + itemName + " - Category is " + catName + " - Quantity is " + strconv.Itoa(qty) + " - Unit Cost is " + strconv.FormatFloat(uCost, 'f', -1, 64))
				}

				var newItemName string
				var newCatName string
				var newCatIndex int
				var newQty int
				var newUCost float64

				fmt.Println("Enter new Item Name. Enter for no change.")
				fmt.Scanln(&newItemName)
				if newItemName == "" {
					fmt.Println("No changes to Item Name made.")
					newItemName = itemName
				} else {
					//item[itemName] = item[newItemName] -> does not rename existing key name
					delete(item, itemName) //is there a way to rename key name other than deleting and inserting new key?
				}

				fmt.Println("Enter new Category. Enter for no change.")
				fmt.Scanln(&newCatName)
				if newCatName == "" {
					fmt.Println("No changes to Item Name made.")
					newCatName = catName
				} else if stringSliceContainCheck(categories, newCatName) == false {
					fmt.Println("This category does not exist!")
					continue
				} else {
					newCatIndex = stringSliceIndexCheck(categories, newCatName)
				}

				fmt.Println("Enter new Quantity. Enter for no change.")
				fmt.Scanln(&newQty)
				if newQty == 0 {
					fmt.Println("No changes to Quantity made.")
					newQty = qty
				} else {
					fmt.Println("Enter new Unit Cost. Enter for no change.")
					fmt.Scanln(&newUCost)
				}

				if newUCost == 0 {
					fmt.Println("No changes to Unit Cost made.")
					newUCost = uCost
				} else {
					item[newItemName] = itemInformation{newCatIndex, newQty, newUCost}
				}

				break
			}

		case 5:
			for {
				fmt.Println("Delete Item.")
				fmt.Println("Enter item name to delete:")
				var itemName string

				fmt.Scanln(&itemName)
				_, okItemName := item[itemName]
				if okItemName == false {
					fmt.Println("Item not found. Nothing to delete!")
					continue
				} else {
					delete(item, itemName)
					fmt.Println("Deleted " + itemName + ".")
				}
				break
			}

		case 6:
			for {
				fmt.Println("Print Current Data.")
				if len(item) == 0 {
					fmt.Println("No data found!")
				} else {
					fmt.Println(item)
				}
				break
			}

		case 7:
			for {
				fmt.Println("Add New Category Name.")
				fmt.Println("What is the New Category Name to add?")

				//var catName string
				var catIndex int

				var newCatName string
				var newCatIndex int

				fmt.Scanln(&newCatName)
				if newCatName == "" {
					fmt.Println("No input found!")
					continue
				} else if stringSliceContainCheck(categories, newCatName) == true {
					catIndex = stringSliceIndexCheck(categories, newCatName)
					fmt.Println(newCatName + " already exist at index " + strconv.Itoa(catIndex) + " !")
					continue
				} else {
					categories = append(categories, newCatName)
					newCatIndex = stringSliceIndexCheck(categories, newCatName)
					fmt.Println("New category: " + newCatName + " added at index " + strconv.Itoa(newCatIndex) + ".")
				}
				break
			}

		case 8:
			for {
				fmt.Println("Modify Category.")
				fmt.Println("What is the Category Name to modify?")

				var catName string
				var catIndex int
				var newCatName string

				fmt.Scanln(&catName)
				if catName == "" {
					fmt.Println("No input found!")
					continue
				} else if stringSliceContainCheck(categories, catName) == false {
					fmt.Println("Category does not exist!")
				} else {
					catIndex = stringSliceIndexCheck(categories, catName)
					fmt.Println("What is the New Category Name?")
					fmt.Scanln(&newCatName)
					if newCatName == "" {
						fmt.Println("No input found!")
						continue
					} else {
						categories[catIndex] = newCatName
						fmt.Println("Category " + catName + " has been modified to " + newCatName + ".")
					}
				}
				break
			}

		case 9:
			for {
				fmt.Println("Delete Category.")
				fmt.Println("What is the Category Name to delete?")
				fmt.Println("Warning: All items associated with this category will be deleted!")
				var catName string
				var catIndex int

				fmt.Scanln(&catName)
				if catName == "" {
					fmt.Println("No input found!")
					continue
				} else if stringSliceContainCheck(categories, catName) == false {
					fmt.Println("Category does not exist!")
				} else {
					catIndex = stringSliceIndexCheck(categories, catName)
					categories = removeSliceIndex(categories, catIndex)

					for k, v := range item {
						if v.categoryOfItem == catIndex {
							delete(item, k)
						} else if v.categoryOfItem > catIndex {
							item[k] = itemInformation{v.categoryOfItem - 1, v.quantityOfItem, v.unitCostOfItem}
							//v.categoryOfItem = v.categoryOfItem - 1 //why is this line not working??
						} else {
							continue
						}
					}
					fmt.Println("Category " + catName + " has been deleted.")
				}
				break
			}

		case 10: //stored lists will be altered when deleting items or categories. is it because we cannot have varying length on the inner slice for each elements?
			for {
				var saveFileIndex string

				fmt.Println("Storing of shopping lists.")
				for i, v := range shoppingList {
					fmt.Println("Shopping List Slot #" + strconv.Itoa(i) + ":")
					fmt.Println(categoriesSaveFile[i])
					fmt.Println(v)
				}
				fmt.Println("Please select a slot to store your current shopping list.")
				fmt.Scanln(&saveFileIndex)
				i, err := strconv.Atoi(saveFileIndex)
				if err != nil {
					fmt.Println("Invalid input.")
					continue
				} else if i < 3 && err == nil {
					shoppingList[i] = item
					categoriesSaveFile[i] = categories
				} else {
					fmt.Println("Please insert 0 to 2 only.")
					continue
				}
				break
			}

		case 11: //stored lists will be altered when deleting items or categories. is it because we cannot have varying length on the inner slice for each elements?
			for {
				var saveFileIndex string
				fmt.Println("Retrieving of shopping lists.")
				for i, v := range shoppingList {
					fmt.Println("Shopping List Slot #" + strconv.Itoa(i) + ":")
					fmt.Println(categoriesSaveFile[i])
					fmt.Println(v)
				}
				fmt.Println("Please select a slot to retrieve shopping list.")
				fmt.Scanln(&saveFileIndex)
				i, err := strconv.Atoi(saveFileIndex)
				if err != nil {
					fmt.Println("Invalid input.")
					continue
				} else if i < 3 && err == nil {
					item = shoppingList[i]
					categories = categoriesSaveFile[i]
				} else {
					fmt.Println("Please insert 0 to 2 only.")
					continue
				}
				break
			}

		default:
			fmt.Println("Invalid Input!")
			continue
		}
	}
}
