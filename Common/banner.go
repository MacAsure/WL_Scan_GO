package Common

import (
	"flag"
	"fmt"
)

type Cmd struct {
	Target  string
	Targets string
	Thread  int
	Output  string
	Dns     bool
}

const Banner = `



██╗    ██╗██╗             ███████╗ ██████╗ █████╗ ███╗   ██╗         ██████╗  ██████╗ 
██║    ██║██║             ██╔════╝██╔════╝██╔══██╗████╗  ██║        ██╔════╝ ██╔═══██╗
██║ █╗ ██║██║             ███████╗██║     ███████║██╔██╗ ██║        ██║  ███╗██║   ██║
██║███╗██║██║             ╚════██║██║     ██╔══██║██║╚██╗██║        ██║   ██║██║   ██║
╚███╔███╔╝███████╗███████╗███████║╚██████╗██║  ██║██║ ╚████║███████╗╚██████╔╝╚██████╔╝
 ╚══╝╚══╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝ ╚═════╝  ╚═════╝ 
																						--code-by iceberg-N
开发人员不承担任何责任,也不对任何滥用或者损坏负责!
此工具为weblogic漏洞检测工具!

                                                                 
`

func Flag(Cmd *Cmd) {
	fmt.Println(Banner)
	flag.StringVar(&Cmd.Target, "u", "", "指定目标url")
	flag.StringVar(&Cmd.Targets, "f", "", "指定文件路径")
	flag.IntVar(&Cmd.Thread, "t", 10, "设置线程,默认10")
	flag.StringVar(&Cmd.Output, "o", "result.txt", "默认保存为result.txt")
	flag.BoolVar(&Cmd.Dns, "d", false, "针对域名目标，提供域名解析可以检测T3协议漏洞")
	flag.Parse()
}
