package types_test

type RecordSuite struct {
	TypeSuite
}

//	record person {
//	  name: string,
//	  age: u32,
//	  friends: list<person>,
//	}
type Person struct {
	Name    string
	Age     uint32
	Friends []Person
}

func (s *RecordSuite) TestMarshalWube() {
	p := Person{
		Name: "John",
		Age:  30,
		Friends: []Person{
			{
				Name: "Alice",
				Age:  25,
			},
		},
	}

	if s.NoError(s.Encoder.Encode(&p)) {
		s.Equal([]byte{
			0x4a,                   // 'J'
			0x6f,                   // 'o'
			0x68,                   // 'h'
			0x6e,                   // 'n'
			0x00,                   // NULL
			0x1e, 0x00, 0x00, 0x00, // p.Age = 30
			0x01, 0x00, 0x00, 0x00, // len(p.Friends) = 1
			0x41,                   // 'A'
			0x6c,                   // 'l'
			0x69,                   // 'i'
			0x63,                   // 'c'
			0x65,                   // 'e'
			0x00,                   // NULL
			0x19, 0x00, 0x00, 0x00, // p.Friends[0].Age = 25
			0x00, 0x00, 0x00, 0x00, // len(p.Friends) = 0
		}, s.b.Bytes())
	}
}

func (s *RecordSuite) TestUnmarshalWube() {
	s.SetBuffer([]byte{
		0x04, 0x00, 0x00, 0x00, // len(p.Name)
		0x4a,                   // 'J'
		0x6f,                   // 'o'
		0x68,                   // 'h'
		0x6e,                   // 'n'
		0x1e, 0x00, 0x00, 0x00, // p.Age = 30
		0x01, 0x00, 0x00, 0x00, // len(p.Friends) = 1
		0x05, 0x00, 0x00, 0x00, // len(p.Frields[0])
		0x41,                   // 'A'
		0x6c,                   // 'l'
		0x69,                   // 'i'
		0x63,                   // 'c'
		0x65,                   // 'e'
		0x19, 0x00, 0x00, 0x00, // p.Friends[0].Age = 25
		0x00, 0x00, 0x00, 0x00, // len(p.Friends) = 0
	})

	var p Person
	if s.NoError(s.Decoder.Decode(&p)) {
		s.Equal(Person{
			Name: "John",
			Age:  30,
			Friends: []Person{
				{
					Name: "Alice",
					Age:  25,
				},
			},
		}, p)
	}
}
