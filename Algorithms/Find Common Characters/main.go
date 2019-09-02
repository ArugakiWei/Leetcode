package main

import "fmt"

func commonChars(A []string) []string {
	s := A[0]
	for _, str := range A {
		if len(str) < len(s) {
			s = str
		}
	}

	minSet := toSet(s)
	strCount := make(map[string]int)
	for _, str := range A {
		set := toSet(str)
		for s, _ := range minSet {
			if _, ok := set[s]; ok {
				strCount[s]++
			}
		}
	}

	resultSet := make(map[string]int)
	for s, count := range strCount {
		if count == len(A) {
			if count, ok := resultSet[string(s[0])]; ok {
				if count < len(s) {
					resultSet[string(s[0])] = len(s)
				}
			} else {
				resultSet[string(s[0])] = len(s)
			}
		}
	}

	var result []string
	for s, count := range resultSet {
		for i := 0; i < count; i++ {
			result = append(result, s)
		}
	}
	return result
}

func toSet(s string) map[string]int {
	destSet := make(map[string]int)
	for _, char := range []byte(s) {
		if _, ok := destSet[string(char)]; ok {
			destSet[string(char)]++
			multi := ""
			for i := 0; i < destSet[string(char)]; i++ {
				multi = multi + string(char)
			}
			destSet[multi] = 1
		} else {
			destSet[string(char)] = 1
		}
	}
	return destSet
}

func main() {
	test := []string{"afcgfbfcdeefhgcaaccgbcehehfbjfgfjacjijibfejjbggbfc","ceiiahidafgdhgcjibjgejhgfifcagiheifdgedidbbgagbiga","hejafbagbagiaheefibiafiigjjdjcfdhcffgiccecgbghcgcd","cdigjbeecehjgjhjgdaedbcddbhjaaidfjdfadibgjhfhahjhj","gcchbbdhcfiefdhcccdbjfhdgcfifhjhefcibdghafhcifajbh","gehjcigdbcbjfehcgdgbeadeejdiaieajhfgadfjfdieecbiie","jfihiccdidbaefbgjadgdgejifjhbaciafhjcdjcgabadhdeai","igbgdcaiicegacfhaijfjgaadafadcdcjhjdidiebfiefbfgic","giibgegaeiciffjgbadbaefdgabfajffbbgccdahiiaccjejji","abigijhajicjahfhchicfhabhgeeagcgiecfbfjhahhhhhbdjf","bceigabhbhcdjdifdaeedgigicffficajhiiggejfceeabgbcb","bcicijaeihjhfgbhihddafhcjgfhgafgeacchaaddjccjbfaah","bcaighaabjcifcgiiehbadiihbabhddfijjafdbebdjgbecafg","dbbjaehcfddgbegbhejccaebacfdaefaaieeghicijjagebejc","cfcfegcbbdiaejebbacbccbecbdeagbdgiigjcddbbhghgijfb","jfdhgecfbbfciffajaeehecdfdbificfabebdbdcjeigeejaej","jdedfebhhgdbbbgbhbfdifcedcgbfjaiiajfcgdbigbaffeeef","ejdfbhhibgdhdbbjdbhbhdbghahfageggdchjfjebbihcgffhg","hgdicjjgbihheeghfdcabhhhbdefaifigjfgeiajagbdibiaec","fbbddgceaigbfhcabahffcjhaiihggjjcgeiedacaiehggeafi","hhdgbbdcdbbjgjgjfggffbedfgjdhjifccdaeihdgahjdcghjf","hjdbbjbjaahjfjjbiajdcfbiiaaeehdfhgaeiihjedddhcjhad","beeafaagajchjdeehgijfaegigfcjidifgcjffhgcghdjieggj","cihafeibdhhcffjcifjhaggjhjbcdejffighcaddiagghegebd","gcebaedchchfefgibgbagfabihcjdhbhbhaccdehjjgbfdgjdd","fdgijfbibbfbagehddbbadhaaeaafjcacaccgbegcfejebbhdj","fgihbhbcfgbjcjbaihjbfajdjgdjgjifhecfbfigafedcfdcgd","egagahjdfcggchabfcdadjgbeggaghaaigggihfhjbgbbaefbe","ccgjebiicbbfcadffgbjhjibceageejibfgdghdececcchabdh","fddhfieefafegchcbihidgjgdjfhcajbgibjhaahefdcihbhdj","fbiajdfebhdbdijjgcacjdgjjafbgebcchigbfhigigdfgdaai","gfjgdfigcfcdjbcdfegeebcgicjdcebdhafecachbebghbhebg","jajbgihaaeabchfgdaijiehajcbcgabcahigdcdfdgejbghhdi","dbijdgfjgbbcihcdijbgdgehhaihibbhhejigbjbafiieedbac","geehiihebaiejcbbadicihgeifficfjaeibahebecjafahefjb","cdecbhibjgahjfhibebijgcbfggidiedcaefdgcjcafgjiefjd","iahghibjaehgghibhdjgjhidbhjdicchgdhcjhabajchbafjaj","ehgejcgajiecifgihafjibbjccigfdafeeigbhbdajdedejfhc","gdaefffedceiadjhefebaigicefjhdaifghbddgfccfafdfddd","gadhiabdbbdeebgchfijhbhhbbhbgcdbaaebegecehigihdihg","hgcdjcifjhhhcabdjbghjfcjiicjidhffhcfcgjgjaidgghijb","aedbebggbegifjdbiiggjbgfhghbgcdfbidfehbgehjhdhcbae","ebicdgciedcghghhcfffdhciagejccadjfdfjdacjafjbfbiff","bjeheechdbaihbacifaegiaiigjecheiaebchegjibdgfhcfac","hjdjcjdbehdejfifediccbfajadeddfijfahbdgfihbcffehic","ggihchfgieidfaabegibbjcibeefegjgiibbaggahghgaaeaag","ghgbbdfgdaaehhjcbiajbiejceeiiaeghhdiefdecjbfifehgc","ehfdicgbffbiegjabdffeeddejbfhfdfjfchbedfdeajibbdcg","cdffadffchiiihbiihccgbbaejafbbdijgabdbfifiajbeaeea","jgddbffdiebeadgchfcbbehihabjjdegfbdbajdacjhhidabbd"}
	fmt.Println(commonChars(test))
}