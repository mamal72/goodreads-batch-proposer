package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const propositionTemplate = `
خوبید %s خانوم.از اینکه بی مقدمه پیام دادم عذر میخام.از آشناییتون خوشبختم. میتونم بپرسم رشتتون چیه و کدوم دانشگاه بودید؟
برخی کتاب هاتون رو دیدم و برام سیر مطالعاتیتون جالب بود. من خودم ترم آخر دکتری مهندسی موادم،تهران. البته ارشد امیرکبیر بودم و بعدش هم شرکت ملی نفت استخدام شدم. میتونم اگر لطف کنید باهاتون بیشتر آشنا بشم؟ من هدفم یک رابطه واقعی و جدیه و مرسی از اینکه این شانسو به من میدید.
وقتی بر مبنای کتاب با کسی آشنا میشی پیدا کردن وجه مشترک قطعا راحتره و من شروع آشنایی در این محیط رو به هر جایی ترجیح میدم. ممنون و مشتاقامه منتظر آشنایی با شما هستم.
`

func main() {
	names := readNamesFromStdin()
	if len(names) == 0 {
		fmt.Println("No name given to generate propositions. Exiting... :(")
		return
	}
	fmt.Printf("Generating propositions for %d person[s].", len(names))
	propositions := generatePropositions(names)
	printPropositions(names, propositions)
}

func readNamesFromStdin() []string {
	names := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a name and hit Enter|Return|↩ key to propose!")
	for scanner.Scan() {
		name := scanner.Text()
		if name == "" {
			break
		}
		fmt.Printf("Got \"%s\"!\nDo it again or hit enter without typing a name to make all propositions!\n", name)
		names = append(names, name)
	}
	return names
}

func generatePropositions(names []string) []string {
	propositions := []string{}
	for _, name := range names {
		upperedName := fmt.Sprintf(
			"%s%s",
			strings.ToUpper(name[:1]),
			name[1:],
		)
		propositions = append(
			propositions,
			fmt.Sprintf(propositionTemplate, upperedName),
		)
	}
	return propositions
}

func printPropositions(names, propositions []string) {
	for index, proposition := range propositions {
		fmt.Printf("\n\n----------\nProposition for %s:\n%s\n----------\n", names[index], proposition)
	}
}
