package main

import ("github.com/gofiber/fiber/v2"
 "github.com/gofiber/fiber/v2/middleware/cors"
 "fmt"
// "strings"
// "bufio"
 "log"
 "bytes"
 json "github.com/goccy/go-json"
 "github.com/chasefleming/elem-go"
 //"github.com/chasefleming/elem-go/attrs"
 //"github.com/chasefleming/elem-go/styles"
 "github.com/yuin/goldmark"
)


func convertMarkdownToHtml(content string, title string) string {
  // We could use the elem go to convert title into a h1 like
//  reader := strings.NewReader(content);
 // scanner := bufio.NewScanner(reader);
  titleHtml := elem.H1(nil, elem.Text(title));
  //fmt.Println(test)
 // for scanner.Scan() {
 //   if (len(line) > 100000) {
    //  var line = strings.Split(scanner.Text(), ",");
 //    return "";
 //   }
 // }
  var buf bytes.Buffer;

  buf.WriteString(titleHtml.Render());

  md := goldmark.New();

  if err := md.Convert([]byte(content), &buf); err != nil {
     log.Fatal(err);
  }
  return buf.String();
}

// Encoding
//	p := Person{Name: "Alice", Age: 25}
//	data, _ := json.Marshal(p)
//	fmt.Println(string(data)) // Output: {"name":"Alice","age":25}
//
//	// Decoding
//	var p2 Person
//	_ = json.Unmarshal(data, &p2)
//	fmt.Println(p2.Name, p2.Age) // Output: Alice 25


func main() {
  app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,   // Use go-json for encoding
		JSONDecoder: json.Unmarshal, // Use go-json for decoding
	});
  app.Use(cors.New());

  routing(app);

  fmt.Println("Listening on port 8080...");
  app.Listen(":8080");
}
