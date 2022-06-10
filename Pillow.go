package Pillow

import (
	cmd "CtrlCmd"
)

var namef = "filepilmoduleforpackage.py"
var nameo = "outtestpilmoduleforpackagego.f"
var impi = "from PIL import Image"

// check installed python3 and module pillow
func init() {
	res := execpy("import PIL\nprint(1)")
	stoppy()
	if res == "" {
		panic("need install 'python3' and 'pip3 install pillow'")
	}
}

// write code - run code - get output
func execpy(code string) string {
	writepy(code)
	startpy()
	res := readpy()
	return res
}

// stop command - delete output files
func stoppy() {
	cmd.Line(cmd.F("rm %v", namef))
	cmd.Line(cmd.F("rm %v", nameo))
}

// get output from file
func readpy() string {
	return cmd.LineGet(cmd.F("cat %v", nameo))
}

// write code in file
func writepy(code string) {
	cmd.Line(cmd.F("echo '%v' > %v", code, namef))
}

// run code and pipe out in file
func startpy() {
	cmd.Line(cmd.F("python3 '%v' > %v", namef, nameo))
}

// show image
func ShowImage(name string) {
	execpy(cmd.F("%v\nImage.open(\"%v\").show()", impi, name))
}

// dropped image
func DropImage(name, newname string, x1, y1, x2, y2 int) {
	execpy(cmd.F("%v\nImage.open(\"%v\").crop((%v,%v,%v,%v)).save(\"%v\")", impi, name, x1, y1, x2, y2, newname))
}


// rotated image
func RotateImage(name, newname string, alpha int) {
	writepy(cmd.F("%v\nImage.open(\"%v\").rotate(%v).save(\"%v\")", impi, name, alpha, newname))
}

// not using jpg
// using png
func DrawRect(newname string, w, h int, color1, color2 string, x1, y1, x2, y2 int) {
	writepy(cmd.F("%v, ImageDraw\nimg = Image.new(\"RGBA\", (%v,%v), \"%v\")\nidraw = ImageDraw.Draw(img).rectangle((%v,%v,%v,%v), fill=\"%v\")\nimg.save(\"%v\")",
		impi, w, h, color1, x1, y1, x2, y2, color2, newname))
}

// resize image
func Resize(name, newname string, w, h int) {
	execpy(cmd.F("%v\nImage.open(\"%v\").resize((%v,%v), Image.ANTIALIAS).save(\"%v\")", impi, name, w, h, newname))
}