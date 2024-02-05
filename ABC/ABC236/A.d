import std.string;
import std.stdio;
import std.conv;
import std.range;

void main(){
    string[] base = split(readln().chomp(), "");
    string[] nums=split(readln().chomp());
    int a = to!int(nums[0])-1;
    int b = to!int(nums[1])-1;

    foreach (i,s; base){
        //writeln(i);
        if (i==a){
            write(base[b]);
        }else if(i==b){
            write(base[a]);
        }else{
            write(s);
        }
    }
    writeln();
}
