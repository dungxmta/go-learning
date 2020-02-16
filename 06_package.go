package main

// ko dung go install =/> import thang theo path
import "testProject/lib"
import libAlias "testProject/lib"

/**
 * cd lib
 * go install =/> build ra lib.a file trong GOPATH/pkg/linux)and64/testProject
 * trong code go co the import "lib"
 */

func main()  {
    lib.TestLib01()
    libAlias.TestLib02()
}