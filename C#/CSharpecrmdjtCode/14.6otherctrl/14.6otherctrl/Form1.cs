using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _14._6otherctrl
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private int a, b, c;

        private void textBox2_Validated(object sender, EventArgs e)
        {

        }

        private void textBox1_Validating(object sender, CancelEventArgs e)
        {
            if (textBox1.Text == "")
            {
                errorProvider1.SetError(textBox1, "不能输入为空");
            }
            else
            {
                errorProvider1.SetError(textBox1, "");
                a = 1;
            }

         }

        private void textBox2_Validating(object sender, CancelEventArgs e)
        {
            if (textBox2.Text=="")
            {
                errorProvider1.SetError(textBox2, "不能为空");

            }
            else
            {
                try
                {
                    int x = Int32.Parse(textBox2.Text);
                    errorProvider2.SetError(textBox2, "");
                    b = 1;
                }
                catch (Exception)
                {
                    // 处理异常
                    errorProvider2.SetError(textBox2, "请输入一个数");

                    
                }
            }
        }

        private void textBox3_Validating(object sender, CancelEventArgs e)
        {
            if (textBox3.Text == "")
            {
                errorProvider3.SetError(textBox3, "不能输入为空");

            }
            else
            {
                errorProvider3.SetError(textBox3, "");
                c = 1;
            }
        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (a+b+c==3)
            {
                MessageBox.Show("数据录入成功", "提示",
                    MessageBoxButtons.OK, MessageBoxIcon.Warning);

            }
        }

        private void button2_Click(object sender, EventArgs e)
        {
            // 清空文本框
            textBox1.Text = "";
            textBox2.Text = "";
            textBox3.Text = "";
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            // 显示帮助
            helpProvider1.SetShowHelp(this, true);

            // 设置时间间隔
            timer1.Interval = 1000;

        }

        private void timer1_Tick(object sender, EventArgs e)
        {
            textBox4.Text = DateTime.Now.ToString();

        }

        private void button3_Click(object sender, EventArgs e)
        {
            if (button3.Text=="开始")
            {
                timer1.Enabled = true;
                button3.Text = "停止";
            }
            else
            {
                timer1.Enabled = false;
                button3.Text = "开始";

            }
        }

        private void button4_Click(object sender, EventArgs e)
        {
            button4.Enabled = false;
            progressBar1.Minimum = 0;
            progressBar1.Maximum = 50;
            progressBar1.Step = 1;
            /*
            for (int i = 0; i < 500; i++)
            {
                progressBar1.PerformStep();
                textBox4.Text = "进度值" + progressBar1.Value.ToString();

            }*/
            timer2.Interval = 1000;
            timer2.Enabled = true;


        }

        private void timer2_Tick(object sender, EventArgs e)
        {
            // 进度条加一
            progressBar1.PerformStep();
        }

        private void textBox2_TextChanged(object sender, EventArgs e)
        {

        }

        private void textBox1_Validated(object sender, EventArgs e)
        {
            

            
        }
    }
}
