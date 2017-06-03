package main

//CreateWaterfall creates a water at origin with attributes
func CreateWaterfall(origin XYZ, o *MCObject) []*Box {
	b := CreateBasin(origin, o)
	b = append(b, CreateSideWall(origin, o, "left")...)
	b = append(b, CreateSideWall(origin, o, "right")...)
	b = append(b, CreateBackWall(origin, o)...)
	b = append(b, CreateBottom(origin, o)...)
	b = append(b, CreateFrontWall(origin, o)...)
	b = append(b, CreateHeater(origin, o)...)
	b = append(b, CreateHeatExchanger(origin, o)...)
	b = append(b, CreateFalls(origin, o)...)

	return b
}

//CreateBasin creates the basin
func CreateBasin(origin XYZ, o *MCObject) []*Box {
	xyz := XYZ{X: origin.X + o.width - 1, Y: origin.Y, Z: origin.Z}
	b1 := NewBox(WithCorner1(origin), WithCorner2(xyz))

	xyz = XYZ{X: origin.X, Y: origin.Y, Z: origin.Z - 1}
	b2 := NewBox(WithCorner1(xyz), WithCorner2(xyz))

	xyz = XYZ{X: origin.X + o.width - 1, Y: origin.Y, Z: origin.Z - 1}
	b3 := NewBox(WithCorner1(xyz), WithCorner2(xyz))

	return append([]*Box{}, b1, b2, b3)
}

//CreateSideWall creates either left or right side wall
func CreateSideWall(origin XYZ, o *MCObject, side string) []*Box {
	var x int
	switch side {
	case "left":
		x = origin.X
	case "right":
		x = origin.X + o.width - 1
	}

	xyz1 := XYZ{X: x, Y: origin.Y, Z: origin.Z - 2}
	xyz2 := XYZ{X: x, Y: origin.Y + o.height, Z: origin.Z - 2}
	b1 := NewBox(WithCorner1(xyz1), WithCorner2(xyz2), WithSurface("minecraft:stone 4"))
	b1.Orient(o.orientation)

	xyz1 = XYZ{X: x, Y: origin.Y + o.height - 3, Z: origin.Z - 4}
	xyz2 = XYZ{X: x, Y: origin.Y + o.height, Z: origin.Z - 3}
	b2 := NewBox(WithCorner1(xyz1), WithCorner2(xyz2), WithSurface("minecraft:stone 4"))
	b2.Orient(o.orientation)

	return append([]*Box{}, b1, b2)
}

//CreateBackWall creates the back wall
func CreateBackWall(origin XYZ, o *MCObject) []*Box {
	xyz1 := XYZ{X: origin.X, Y: origin.Y + o.height, Z: origin.Z - 4}
	xyz2 := XYZ{X: origin.X + o.width - 2, Y: origin.Y + o.height - 3, Z: origin.Z - 2}
	b := NewBox(WithCorner1(xyz1), WithCorner2(xyz2), WithSurface("minecraft:stone 4"))
	b.Orient(o.orientation)

	return append([]*Box{}, b)
}

//CreateBottom creates the bottom of the falls
func CreateBottom(origin XYZ, o *MCObject) []*Box {
	xyz1 := XYZ{X: origin.X + 1, Y: origin.Y + o.height - 3, Z: origin.Z - 3}
	xyz2 := XYZ{X: origin.X + o.width - 2, Y: origin.Y + o.height - 3, Z: origin.Z - 3}
	b := NewBox(WithCorner1(xyz1), WithCorner2(xyz2), WithSurface("minecraft:stone 4"))
	b.Orient(o.orientation)

	return append([]*Box{}, b)
}

//CreateFrontWall creates the front wall for the water to cascade down
func CreateFrontWall(origin XYZ, o *MCObject) []*Box {
	xyz1 := XYZ{X: origin.X + 1, Y: origin.Y, Z: origin.Z - 2}
	xyz2 := XYZ{X: origin.X + o.width - 2, Y: origin.Y + o.height - 1, Z: origin.Z - 2}
	b := NewBox(WithCorner1(xyz1), WithCorner2(xyz2))
	b.Orient(o.orientation)

	return append([]*Box{}, b)
}

//CreateHeater lava is needed to prevent freezing
func CreateHeater(origin XYZ, o *MCObject) []*Box {
	xyz1 := XYZ{X: origin.X + 1, Y: origin.Y + o.height - 2, Z: origin.Z - 3}
	xyz2 := XYZ{X: origin.X + o.width - 2, Y: origin.Y + o.height - 2, Z: origin.Z - 3}
	b := NewBox(WithCorner1(xyz1), WithCorner2(xyz2), WithSurface("minecraft:flowing_lava"))
	b.Orient(o.orientation)

	return append([]*Box{}, b)
}

//CreateHeatExchanger protects the lava from the water
func CreateHeatExchanger(origin XYZ, o *MCObject) []*Box {
	xyz1 := XYZ{X: origin.X + 1, Y: origin.Y + o.height - 1, Z: origin.Z - 3}
	xyz2 := XYZ{X: origin.X + o.width - 2, Y: origin.Y + o.height - 1, Z: origin.Z - 3}
	b := NewBox(WithCorner1(xyz1), WithCorner2(xyz2), WithSurface("minecraft:glass"))
	b.Orient(o.orientation)

	return append([]*Box{}, b)
}

//CreateFalls creates either the lava or water falls
func CreateFalls(origin XYZ, o *MCObject) []*Box {
	var surface string
	switch o.oType {
	case "waterfall":
		surface = "minecraft:flowing_water"
	case "lavafall":
		surface = "minecraft:flowing_lava"
	}

	xyz1 := XYZ{X: origin.X + 1, Y: origin.Y + o.height, Z: origin.Z - 3}
	xyz2 := XYZ{X: origin.X + o.width - 2, Y: origin.Y + o.height, Z: origin.Z - 3}
	b := NewBox(WithCorner1(xyz1), WithCorner2(xyz2), WithSurface(surface))
	b.Orient(o.orientation)

	return append([]*Box{}, b)
}
