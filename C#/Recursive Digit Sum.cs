using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Collections;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using System.Text.RegularExpressions;
using System.Text;
using System;

class Solution {  
    private static string createCompleteString(string n, int k)
    {
        int nLength = n.Length;
        Int64 sum = 0;
        for (int i = 0; i < nLength; i++)
        {
            sum += k * (n[i] - '0');
        }
        return sum.ToString();
    }

    private static int CalculateSuperDigit(string t)
    {
        int tLen = t.Length;
        if (tLen == 1)
        {
            return Convert.ToInt32(t);
        }

        int sum = 0;
        for (int i = 0; i < tLen; i++)
        {
            sum += Convert.ToInt32(t[i]-'0');
        }
        return CalculateSuperDigit(sum.ToString());
    }
    // Complete the superDigit function below.
    static int superDigit(string n, int k)
    {
        string finalString = createCompleteString(n, k);

        return CalculateSuperDigit(finalString);
    }

    static void Main(string[] args) {
        TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        string[] nk = Console.ReadLine().Split(' ');

        string n = nk[0];

        int k = Convert.ToInt32(nk[1]);

        int result = superDigit(n, k);

        textWriter.WriteLine(result);

        textWriter.Flush();
        textWriter.Close();
    }
}
