package main

import (
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	_"github.com/BurntSushi/xgbutil/xevent"
	_"github.com/BurntSushi/xgbutil/xwindow"
	// "github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgb/xproto"
	"strings"
	"time"
	_"reflect"
	_"image/png"
	"os"
	"encoding/json"
	// "image"
	_"strconv"
	// "runtime/pprof"
)


// type TablesAll struct{
// 	Tables []Table
// }

type TargetTable struct{
	Name string
	Big_blind int
	Max_players int
	Max_buyin int
	Chat_x1 int
	Chat_y1 int
	Chat_x2 int
	Chat_y2 int
}

type TableWindow struct{
	Window_name string
	window_id xproto.Window
	current_chat string
}

func main(){

	// start_time := time.Now()

	conf_file,err := os.Open("./conf.json")
	if err!= nil{
		fmt.Println("error while OPENING json : ",err)
		return
	}
	decoder := json.NewDecoder(conf_file)
	playing_table := TargetTable{}

	err = decoder.Decode(&playing_table)
	if err!= nil{
		fmt.Println("error while DECODING json : ",err)
		return
	}
	conf_file.Close()

	// target_window := "Chandigarh Crack"
  X,err := xgbutil.NewConn()
  for {

  	if err != nil{
  		return
  	}
  	// fmt.Println(X)

  	all_tables := make([]TableWindow,0)
  	// fmt.Println("len is : ", len(tables))
  	// fmt.Println("cap is : ", cap(tables))

  	window_ids,_ := ewmh.ClientListGet(X)
  	// fmt.Println(reflect.TypeOf(window_ids))

  	// adda_chat_rect := image.Rect(playing_table.Chat_x1,playing_table.Chat_y1+20,playing_table.Chat_x2,playing_table.Chat_y2)
  	// count_tables := 0

  	for _,win := range(window_ids){
  		name,_ := ewmh.WmNameGet(X,win)
  		if name != ""{
  			// if strings.Contains(name,target_window){
  			if strings.Contains(strings.ToLower(name),playing_table.Name){
  				add_table := TableWindow{}
  				add_table.Window_name = name
  				add_table.window_id = win


  				// dis_file2,err := os.Create("chat_"+strconv.Itoa(win_index)+".png")
  				// if err != nil{
  				// 	fmt.Println(err)
  				// }
  				// png.Encode(dis_file2,chat_adda)
  				// dis_file2.Close()

  				// fmt.Println(name)
  				// fmt.Println(win)
  				all_tables = append(all_tables,[]TableWindow{add_table}...)
  			}

  		}
  	}



  	fmt.Println("tables : ",all_tables)

  	// adda_chat_rect := image.Rect(playing_table.Chat_x1,playing_table.Chat_y1,playing_table.Chat_x2,playing_table.Chat_y2)
  	for _,table := range(all_tables){
  		fmt.Println(table.window_id)
  		// window_full,_ := xgraphics.NewDrawable(X,xproto.Drawable(table.window_id))
  		// chat_adda := window_full.SubImage(adda_chat_rect)
  		// go ParseImage(chat_adda)
  		// go ParseImage(table.window_id)
  	}
    time.Sleep(1000*time.Millisecond)
  	// fmt.Println(time.Since(start_time))
  	// time.Sleep(5000*time.Millisecond)
  }


}
