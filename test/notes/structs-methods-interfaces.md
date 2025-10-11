### Now that you have some understanding of structs we can introduce "table driven tests".

# Table driven tests are useful when you want to build a list of test cases that can be tested in the same manner.
[ 
    func TestArea(t *testing.T) {
areaTests := []struct {
               shape Shape
                   want float64
           }{
               {Rectangle{12, 6}, 72.0},
                   {Circle{10}, 314.1592653589793},
           }

           for _, tt := range areaTests {
got := tt.shape.Area()
         if got != tt.want {
             t.Errorf("got %g want %g", got, tt.want)
         }
           }
    }
]

# #### The only new syntax here is creating an "anonymous struct", areaTests. We are
# #### declaring a slice of structs by using []struct with two fields, the shape and the
# #### want. Then we fill the slice with cases.
# #### We then iterate over them just like we do any other slice, using the struct fields
# #### to run our tests.
# #### You can see how it would be very easy for a developer to introduce a new shape,
# #### implement Area and then add it to the test cases. In addition, if a bug is found
# #### with Area it is very easy to add a new test case to exercise it before fixing it.


# %#v is the format string to print struct. 
