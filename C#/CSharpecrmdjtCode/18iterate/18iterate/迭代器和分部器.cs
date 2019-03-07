using System;
using System.Collections;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _18iterate
{
    public partial class 迭代器和分部器 : Form
    {
        public 迭代器和分部器()
        {
            InitializeComponent();
        }

        private void 迭代器和分部器_Load(object sender, EventArgs e)
        {
            Family myfamily = new Family();
            foreach (string item in myfamily)
            {
                richTextBox1.Text += item + "\n";
            }

            // 设置下拉框
            comboBox1.SelectedIndex = 0;
            comboBox1.DropDownStyle = ComboBoxStyle.DropDownList;


        }

        private void button1_Click(object sender, EventArgs e)
        {
            try
            {
                // 实例化对象
                account a = new account();
                // 获得数据
                int M = int.Parse(textBox1.Text.Trim());
                int N = int.Parse(textBox2.Text.Trim());
                string op = comboBox1.Text;
                switch (op)
                {
                    case "加":
                        textBox3.Text = a.add(M, N).ToString();
                        break;
                    case "减":
                        textBox3.Text = a.sub(M, N).ToString();
                        break;
                    case "乘":
                        textBox3.Text = a.mul(M, N).ToString();
                        break;
                    case "除":
                        textBox3.Text = a.div(M, N).ToString();
                        break;
                    
                    default:
                        MessageBox.Show("没有这个选项");
                        break;
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message.ToString());
                
            }
        }
    }

    public class Family : System.Collections.IEnumerable
    {
        // 创建一个string类型数组用来存储家庭成员
        string[] MyFamily = { "父亲", "母亲", "弟弟", "姐姐" };
        // 实现接口
        public IEnumerator GetEnumerator()
        {
            for (int i = 0; i < MyFamily.Length; i++)
            {
                // 依次返回每个元素
                yield return MyFamily[i];
            }
            
        }
    }

    // 分部器的使用
    public partial class account
    {
        public int add(int a,int b)
        {
            return a + b;
        }
    }
    public partial class account
    {

        public int mul(int a,int b)
        {
            return a * b;
        }
    }
    public partial class account
    {
        public int sub(int a,int b)
        {
            return a - b;
        }
    }
    public partial class account
    {
        public int div(int a,int b)
        {
            return a / b;
        }
    }



}
