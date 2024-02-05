import std.string;
import std.stdio;
import std.conv;
import std.range;
import std.algorithm;

void main(){
    readln();
    string[] takahashi_plates=split(readln().chomp(), " ");
    string[] all_plates =split(readln().chomp(), " ");
    int[] prices = split(readln().chomp(), " ").map!(e=>to!int(e)).array;
    int except_price=prices[0];
    prices=prices[1..$];
    int sum =0;

    foreach (plate; takahashi_plates){
        // plate: 高橋くんの食べた皿の色
        int current_price=0; //食べた皿の値段
        foreach (key, color; all_plates){
            if (color==plate){
                current_price=prices[key];
            }
        }
        if( current_price==0){
            current_price=except_price;
        }

        sum+=current_price;
    }
    writeln(sum);
}
