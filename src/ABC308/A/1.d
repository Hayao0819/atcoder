import std.string;
import std.stdio;
import std.conv;
import std.range;
import std.algorithm;

void main(){
    int[] nums=split(readln().chomp(), " ").map!(e=>to!int(e)).array;
    int previous=0;
    foreach (n; nums){
        if ( previous > n || !(n >=100 && n<=675) || !(n % 25 == 0)){
            writeln("No");
            return;
        }
        previous=n;
    }
    writeln("Yes");
    return;
}
