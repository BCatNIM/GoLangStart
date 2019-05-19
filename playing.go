package main

    import "fmt"

func main () {
    const (
        lightspeed = 299792 //km/s
        hours_per_day = 24
    ) //end of constants

    var (
        distance = 56000000 //km
        speed = 100800

    ) //end of VARs

//	fmt.Println("Testing Go")
    fmt.Print("My weight on the surface of Mars would be ")
    fmt.Print(203.0 * 0.3783)
    fmt.Print(" lbs, and I would be ")
    fmt.Print(41 * 365/687)
    fmt.Println(" years old.")

    fmt.Println("---")
    fmt.Println(distance/lightspeed, " seconds.")
    fmt.Println(distance/speed/hours_per_day, " days to get there.")
    fmt.Printf("%60v\n", "--- Listing 2.2 ---")
    fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
    fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
    fmt.Printf("%-15v $%4v\n", "Mikayla Air", 1221)

    fmt.Println("")
    fmt.Printf("%60v\n", "--- Listing 2.5 ---")
    var age = 43
    age = age + 1
    age += 1
    age ++
    fmt.Println(age, " is my age")
} //end of main func
