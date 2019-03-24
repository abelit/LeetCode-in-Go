package problem0955

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []string
	ans int
}{

	{
		[]string{"nobhttjnanwsfjaoythqlziymsvpqgxsfvpzmgoloppabwuzxyxrtvleozeffvuiobsnwrejqzaxrapmgcvnguugpqwaktbjuowr", "vnzaupgircggfdazttarjooyggjjfwgoezdesayffmfpdyfqzknbvkeunsxgfkvilmqzqxyxymbfklqowvqshwgrhqqutovxdquj", "jrxxdrlmeyfcsmsaxntrbxjjhsjitgdfsjdsqmzkaajhupidktqiwopbthaaukvgpzphgwsrzobmolgamoegiprphpnuyswvmrer", "rvqsgytjnymsuhpthsfixvsmaaffmjzczrcaviodivzxzjafdwzaftnmfneopzfucbvvwzzzorbpdnbxynsazmtgszimvmnlkzei", "qwcuafxeawbekrgkqzqjdjlygxasmmydkrpidnxezawcblqhdcldymxxytpllanuylgzdfsporccczayczwolxszolqgouaftebw", "hgacebtnnorodnjbwzfrxodcoexgcoljheypkswucrdwuzykrzttviuyooirscmazliqkzqtjucgfoewzlrufmuipnzihmrewrgp", "poiiqlckvgqjguiloqsambwtazngpkbxlxuyudpazegfjovdhcbexqyuzfphoaqqxgdqncqukxcnqzxqrnztogafefllnooldvzd", "saxjgvpdzvducqohlhviviigqwnvdfbjwfzudrskobxoxcnwkgkgkljwayazrdvbbbguhskcuhcpqrmgjqtwnsxubzxzcdyrfurh", "nlyhfleljrgkvcfrttbhxhsvqlehskvfkemzqlptfanctrgazyacjchfrwmxjbkdezpgyhkbtecputkbelracwgwqsmzumtnqazx", "xyfdjkzugkhliomgpmwxewrqqptenajhwmlnahepjzjnzybqlnkzzolvnvvijdnavmykpygksocyxqcmvargirtngaeubransrjz", "hmauiygfrvtallotnwhtweykggifyocnygxbibbishkwkshkbkobhwfohssxdlihknzzlcbfafddiqqznhqlgeycpzpdzkgwlqpl", "dqnorvruwcmrprqfucwtxqwjdtxonctgoypyqlhfkqsuvtbmfcfrbjjdglenteyzkksqcwhkgxdgafzgnecforlmsffpvqwhciaw", "dkcdkpttoxajqnbihprlrdkwwiivuyvxlbcdnfeclaqmiprvptahupvddgphhunjwmremymsnudpwnxdaovvfavxdmgfaotqbijx", "ojgssedmspxmftngpnicmysghydvlbfsxzmayhtjcqtvudnrlffnpuvlslcihxexqgpvrzcamzdsouzupaxhafjogryatvnsrmgt", "xqrjxzgowfzphmolckctnslrdmpvymmjrkmcabqnzebkbsqjuzifmvbuthdvpfkrqglldsyrmedwkyxzqjpggtwhiguefekcgqrz", "umvcxsfbqfypamhfqrhrifdtbkxpfjwfrmnfxdvghwgmlqhlotddlkgdbktzzmxyzbrvexwxwxdzzjynoojcuuyfljdwsnjupouj", "jxzhumhaecvsnolmzyjwxdxuvmcfnarfpqobjvvrguzekmwdgrqqwfwjdabcmbdgphzkkvckxbejkbgnphtiwtuwympkgqzyevaz", "ryrkzfjrvphycmhqovxdodjzcfhfisbsyxcxhgrolbqqsjmzpndgxjlpyiiarxbspnshcjjcfkemfjdkrtzmtbypymrpwckgksax", "xbnhcitibihydssekxkuxqxxnsypykwejmpbpfaajpippykfupcanqsrvuzjeqlwyxfkbhmvdteunxpecjduyiwdzjrwirtgffhr", "hlznpdalynzlzpvhwczfssdvblgdujzqmdietyrghrvfrlmtjvpjgxcycmzsucbbwlthbrsqbrewpqyfskitalevfsuyytdkvgfq", "zxmetstwhjgqbkcnuaqhtjilzwwuzuwvwcbeoqfvzohgbnsqadazmupycngvsqrtforvebxlhzfljbxldpjuprpsmjzfrcrdmeaq", "kbulkjgdfokeybdomgseeyblnhwlsptzvhygakzbzjptnnyhtlhgdmzfepsokakmcaeeytlhqyfquxpcktxqlxpfhfehksocgvfv", "zzzngsbzpamxbgkmedczejbaaafjvtisepbkrvstsabncxsthfjfioizxvppfqwazaaknyhxhufvvywusfqiblokjpqbeqmawilx", "opjagfdtqsgsodzbidjimyblcpqivrmcrfaulbuzulfwavimkmjrjfzbeuvtdpphspptrncuyefxsoctgrbziicpenckxitdsxev", "bxgrnjveicjmtgzthacrmmhhbhapmyakfkfbmldtvjdsuqyydkkwstaekeaxbazxleaapkbrzhfyonrqeuxeqgyreqhelkhvxebt", "wppdwufbsdcwxqufibzydwspjtjoygkhckvbnhtkaxcarrilgggdjdpqxaoopswvwnmqpnjlvqfyopsdemdblrkhkjhubmqpcnqa", "mkyrpfeyouhbcrrgrjtcxwafzpkghjcnelwikbbgktybexkyhbmkvplwpebormqhhdszdcmzdwgmvhmmimyfrqfhomcbjsxvgdrg", "jbpbchsjzbfrptibifazielpzrrosxxrkcrnyimmetmjejmkneihvsvphkutryrwsihvgdfkufiagtilhzxyjdiygqbzanjqrhfx", "cjzlxpnzfsmqzrbgkwdxtozaotwhajsngxqztkzkvzfkcjfcjnbztydhwdqvynpbtvhywocjflicratlgvgmojwzonpczjbybynm", "irudzaapcaakadlchpwuringeavwxzpnvkkefpovuhomjkfkbktmytfjcrnqaptkbacoozbbhuiqsklagfyksrxveetmokxtsaeq", "ohirkpimrbaickgzbejrbxmcsxpwydchtdwhiiyqiaiohqwolhhbvafkkneowxmaqrhaalyaffjxreazftdtnyylobbbxgswskew", "uqbbuhgywtkoqyuaygegxoeowczivecgkciabypgfgplmqowmojqtdknonxqkfyhfwomzdirbvkukmdequrezaqtuojmicaxcpfk", "cpduibajwldtjzkbigjxrpfhyzjvoewookzgnfuwhamrrypfsjzzquharpcevjzsbljixfiupwljcjxrtrrftiepsgoruykgcior", "saddyfxdpvctiotakglrdbaowduvvupqdwkmnealwxgpntvtdsfmriessrwwidbwratmebfreqlnahcphjvionmidxuzwjziggsr", "hhdpgzrbidmwzsswbnjzaiudxxkbbbuoezszvyqkdcucgahuwzdvghlhdnevuslzhdhspqpxqilpigojohzlrokaqkjpawhylpzk", "nrqlfenyomdhmxqvtqjfcphqfqjtduxijencdvudyqqhtygaziylucuuvkuciwrmttagrnfojslwgchdogdoupsbsbglhdzdufnn", "dhbgukrrrnioxivpfciaseovghiiqfiadvmojeyzdbszsoumehowbxhwnyvuxhyqgzgygfflvflyofodjoittschlurcdamlgjww", "calhlfhgppsikxpcpkjjoqjwivrdpsvwtmhgrjonnbzcyvmqzwxcddyixururvrbpnpbfjiadflzhvycgoivncoteplodztiugon", "oeoayqpytrbqogrpjqrftyuwrnmxqtzrudzgvbjedaguokeighusaoomgevpywragjgrlixkanmgzctgsdhkbxjsfbpodtoqisfe", "qcyaosryeclcqsyywgbmmnmtudytqbiwneemyyswssgjnasxyhsmkbjbjrpgtrruymyjewfwkomlzusvihpxuettmhjraroyklrx", "zixpuufjugkgaardcudctjkijaehtstncoircazqxhvbcfughqhlwatxbyfmeugbmseivolistmmkzhdtkfebexdusgjszwamcvx", "twiqzwqxfsgdhkdwjalejfzxudwvqpclxzqwzmfkxllzunmfmunnbjnfpzxkznmxwefyfvnhhrngiiagzvwgxlureiyhbteycsnl", "aibbqdrtsvhhnrqofawwdbqmphalczvyjlrotqkhttjnifbqkwotpphnjfyikfvssmgsnhkgjinkypzrdlxfsohmwnoynbehgcxz", "pxvfwtwlafycvqzozebolsompldxfhmnomnjddakhmdysyjmlalmkmmpcshvsclpfbmhviksbrogjrdooixctqlgalayioezshxo", "obzvnqimmvyxxjgnysrowbypfoolfkygbosbvroecunahoxtqtlvxweinutubcuaxmnpterplxokvbykabpmgfandumwsmqfiovj", "vdvvoukuisnggrahehjvmcffjqssorszpgvlyqkwdxfkfcyrpdyiihrzngvsolacrqluphbsgbolaxwqyxuhgrkzkkyacxxrcwtq", "xfvsfetyfzidgkcvraplasyzspypkpyeidpmnxodwjzrsvsxzbfmoemyvvzhtbsnnbmztaiijcpapxukeqcyjokvvzlwpfnjluiz", "jcndqdmqyrzphvhmzqwrfqbujzcovpprrvamuvtxpmacdppfpstjcjqehvlhkhlnrmcqdcjgdppdbziqczwtrnuxlahtljtbzjgn", "juotqnaeqsispfezjmkvzroivnnczpzxwnwbzodsyzjpqvieylmyiqejdsiloeatycacrcpgaapkpeafuoorjrymrqevfaqfczpa", "iblikloirsywhdmnfjwwozryvcwogrihgqfjqyfemaasfbosqielhseuplcxpvelercddvnyhcpqvlpjwqcbgnsmmlbusgirlzzb", "pezrxkgtsgoyxlfncnviteszrjcympupxestdeiirjamnhkulmcsubjonvjmhplunjtppalypoqfxvewhqxoiexlzgdcpanxfzup", "qmrzspkklimygtkpfdwjhnahhfndliodqpjmnftwqlqaaqgqvcyupljhrowivjylzyirftjvonqhitnzyaigntuijzkxtncqhksm", "juvuznetreylakribucjywjsckpiggnlwrjffabzjmuegdvxzyrjlclitllsayxwquagfgogtyqqgotntbxyobemvskwwpfchskz", "sxpusrsqduguxjzcmaaemqifazwmohbtfepzzwwsysmmsfbnihhluiahkoyhszcowjpfzzrizfrilpicxgiibipamzyznjhicfdi", "seyqpopooecldhydmqambxnxxiachmfdvcfefslfgxneownavqvzktfhqlldcvkopcvowduundrmurwcvtavkujkgsjvxqnmvoio", "xuqwkgeaugprlrqvxbynubpqxmrtzbgskkvhmmftyqxwksipefywfwbzkbyndufzrrnijqhhhasjagkunaanexpnrkxvwnmfneaf", "rsprllqjicbujdedunyeaasiwkvxeehszjjqwaqtlgyyoslilvomismzxzmtsnqdxnhfshzpuysjtwacwrfgxfvetcxribrshtbr", "onixzpwuxaelchmzvzzhezijjwyvwcesizwaarvlubjobxemlrpiosxtuocfbpnacbnrsuvhekskskqbrrmvomzrcjezunmzchlj", "xatqvgufdguhenbdqxhqyqlnsoflukpcsybpqkhqmvhamxnrzohejfddwhdhotwexbqppnlksgsuvqfgwwpouvgqcmmlxbksqgjd", "uvgcwzkmnrazrvfyffwsguooygpmqwsurgvxfjnfrywwfczmikhfymuyluflwchxsfhkmbbzslsvlhicfxomhmmohcfdmzkaahbm", "qgmqgzqkmmwuoveyvjcffebidcwvldmoszgzzrwdynzdmjuqgvfnaiykdhtonssyekywpmbeuftijgdvugvhvmurdhblohmvjxbd", "vyjiooeamfklryjkaofokbcurltjujrnniltrfrafgcmowwyspylckvbaoawahsdkdtsfsgzixtmwsqqylfddjcobbhiudsuermr", "istwcyruifilmalcynnpseuvdapbcrsowblopftdufiwoqytbzddfeuvcfkjcgmaoqmaqjkakatpronphslvsyhpcttwqqgdydcb", "onuapbtebwhjjxclktndnguegxdtenpipndduzavsxmhetykqqdbksvtdskmjojpjbzbgpawlvtrnxrmfhsttihyhgzhskgzfhxa", "sbtceulryccnzzuyzhmasuxpvibykoealyylaaxyxuzfqduppjjrjpbgyhkkitrpetncltrnapuaewnwssroctfbkuigrlutetjj", "uvfgtiaeieiybeslslxpzqtnrqteooxsrxgfoihuefuwwqdfoxxaneqfnsqzvdpbilvyyhmaeauewaolcambienmubdanhkuxcge", "ensmqnrzjltbdtyiaxzxzegyxpoypskxtunahevcxfqkcpbytdsysndgctalnlijhhrxbcgqviuhwbprllxciruyijvxxhpzscgo", "dynidnxljmbijllwwvjqcjoqgalikgzomxmlvoveyzshfvotjnbmdbyeipopmvlrbassfdvuqrukwjaxqgwxlqrcyxueykznnxjz", "rkxuubhscwjocyelteklgxifnorblunyadrjidsefscqznkwmhxkmztemaiajkciromvyfzzwiunyjrzmlwmldlhfthajvklatvl", "cgnlgaqrpqgktrylnthagdmsnhkggvyaqckfincsnavmdtuwjsrjrrskfrslrgxqcklgbwtlxduttvqttikfjrbskmwjqtbyavmo", "ewzflfemetkqqrnsnusbloyqwyvktkrhysxypsniiynivognelvujwcrlgjxiorglayuglbuiyuuydukdjrhzjiafupdibkcbgrk", "wjqtoxsvljfesolvkdyfqxmmpzvoocnyvxhwjpwvlhjpienyepflpnngzmvmztmgjawqegcvkzuyacmqipwulphnjaboldifpcoh", "tmckqvfumjtpmjmijytbuhrnnfltyqjagozieqgswyfpjgudduipqnfogorfugtjojmjdsautwvhivsylevfoznwnwbrjfknrrxg", "jhzqwbukqsbqfhnbvfftreqqjkzxbtvxfkeaxxccohsesdnttfgasqyiqsszwvdsjwmldnxvrwvyzeeagtjxxhklquyfruwurbsi", "tzuwyyojswwldxpzzhlylvhodsxirmwqklsmedegmbnzatchzlivdzzcqnmctclyxaiwmkinzuwjiehhjblrcevtpvjxzpkrkprb", "lsfzutypbznpiuplegezeouhltefsmqafycudgumetkmqqqmccxiuplmreegatxzsxxosmiepdwougsglifctzeifnvgiikuerea", "fwnmtllbophqgwikkmglnxvcftseyyuthnsnzazyptrqmaoukmuqkfrhzegnzpcowkzyyozmlqwqprgmhlgmiqxifbjvjyfvhtmp", "esnqustifizyijbngkwmkeworaathlqkenmdacnrcsdplkeipdjjditjofjmaylerlankqxelywqshookugwrcedzafxheuhpich", "pchwcigplovuhyqfdljqvmixrzinzmjkwzbshtcrupjcqejyrzeipibuqyiexfhxhsgvvmaqgtwwnmuclhvfnjsfozxcndelgyib", "zxrjngumnjhzdavthekeadsugmjnanyecxinulhcpfkynyncavbyeuyccyxpcxgxotjpjxfdgaxcwerybtpuhxgofircivkkjqzd", "yuaobmfqiunwovzyulkkhsrsarrauykfdkkcipjojfjmarlzhahcwwxdhknoybmxyinlivwupsxemzqrwxekbbogjsfebjucoxxq", "ammhctpsepgbagoafnpwaumammlcjcfbxqmztvzjumwujeylpotaswwtgwnkacngvsyikffefdxgcnvvqbayyrjvudrtgdhpuglj", "edknamctunqcijhnfcdiuseuvopwhkiwxcxdepstzesfyobvkmnrmfxgsybpdotwjtjvakwvhzximtiskkwzydehsthhyplbnhrc", "sifrawqfgcgbykecxppsiwxszqumdtxquguynqmevamwirhsdxvwfuytaljabhzrqoybwiyvywxmquwjboqqkprdxwblzzgsrqvi", "zopzxtlxsytnycerbtlbsokjadatritemzbvwmxxfdkyzamlezqekudwprprnlnszmfapweczixmzwuzgsbnkktlogfvvsajuwoe", "kfknjfcnqletleexwhwaiigsztoqczbpopsjcshjjoqmvposltoshvdtrxnoalfvvptmaaszxexyycnsivjdrwijzkzhpyucdfik", "yebnuukkmsldxswocmqhjsrmpymyeiyjyirbxfzpwjdqpbtiaaqnpuyifurtvjhieudsymmnpxydasmgitzzzcgnctyckzsohlwm", "dymvgjducexyedlbtponapnebmdkfdogouqpyibhuhsfprntcxxzykawvrprjzceaydjlozevpyeaxvckuqmxxjegfkvmxouyukz", "ulkhumuqnrxeusndjcjqxqjaivcsvktgdlxbmyydijrugqwhckxdghxanirqlzmufpnhnmkdsiymuqurtoodmqotufmyrgkvvlpl", "sxzkgmtehoqeouqbaeopsfratiekrelbfpqyrfhraxbkkiqmusrqcdtswbfnjndeehaawzlgayynzqywtudugdwyeupnraptwjjf", "hoxexviivuqqmggfxhsqxjmaqqnbunovmrbcakkhdqjkpopkhdaylpfrezzeckfgbijdunxpdzyveansmqpnfmydczzpmxtlxkbp", "xywmzjvuexvcnrbjeuaxabecistwssxeqxhldvpmbyehiyampuzqpagjebgmtetjmmcjncioegyvvilzrcwphkhhfxcxhnmuukwq", "avgqdewdnczmhjlxdvuejphbjjtqplsbrrunsknsushattrwbgugugeqembvahdkxhwvphkwtnzcfycoclwwzctcubmibztupbgb", "fowggwthyyvgdheaogzhlyiysnzujudxmfjvuwspqqdwdgzlxljursdsesyvfasqtmtzmohjehzfafilrlwwvchfwksqyrwizakx", "eubkidkyoqikppqlwealwdqkhimuvzuqnjbkjhwrojiurpukonttxnmkwlwtvujpyzxcyxutugzmjlpylrliiwbihsjsanaazofh", "jhagizqiwepckxdzxufccslxizelsgmmzzxymouzttrxoeshxqypqzrbuthsbqbnfczlnecxeczqgoxsbsleitmwpofwaxwvkmmn", "ophsohylthngescjeddixbgmtrhvnysmmylucbgwlpdqnljsztebfqcahjeqeeozsssfuisualzrgbihpdckrqarsylkklmnevpg", "czhvbrbjropeoagempmmsicdmtqimvtbkmndujqcgfqidsnvqphffdhwdoelmwiustgnadcgpzzrluxgfrjcsqmurdlzkactiqvm", "xddmrmupfsnqlidwsxopxbpbmkiklhqzmfxkpvfuebvwnkrjhmpxxmbrkplihiabctzoqavsilzsmeoujlbhcnclaohphekvwuhq", "cchrefbftixmfiekisqezsbvlcrtiihlgjuahdocpasplultieammaqmzqaqrefakecyvuhwiyztzemkbezpjbpocxirqwcyzjdj"},
		74,
	},

	{
		[]string{"abx", "agz", "bgc", "bfc"},
		1,
	},

	{
		[]string{"ca", "bb", "ac"},
		1,
	},

	{
		[]string{"xga", "xfb", "yfa"},
		1,
	},

	{
		[]string{"xc", "yb", "za"},
		0,
	},

	{
		[]string{"zyx", "wvu", "tsr"},
		3,
	},

	// 可以有多个 testcase
}

func Test_minDeletionSize(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minDeletionSize(tc.A), "输入:%v", tc)
	}
}

func Benchmark_minDeletionSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDeletionSize(tc.A)
		}
	}
}