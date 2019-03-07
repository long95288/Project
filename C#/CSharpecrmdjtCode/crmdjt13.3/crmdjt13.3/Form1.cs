using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace crmdjt13._3
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
            label1.Text = "新的文本显示";
            // 可见
            label1.Visible = true;

            // 不可见
            // label1.Visible = false;

        }

        private void Form1_Load(object sender, EventArgs e)
        {
            // 设置接收enter按键
            this.AcceptButton = button2;
            this.CancelButton = button3;

            // 设置只读文本框
            textBox1.ReadOnly = true;

            // 自定义设置密码框
            textBox2.PasswordChar = '&';

            // 使用默认密码框
            textBox3.UseSystemPasswordChar = true;

            // 设置多行文本框
            textBox4.Multiline = true;
            textBox4.Text = "昨夜西风凋碧树，夜晚上西楼";
            textBox4.Height = 100;
            textBox4.SelectionStart = 1;
            textBox4.SelectionLength = 4;

            /* 富文本框*/
            // 多行显示
            richTextBox1.Multiline = true;
            richTextBox1.Text ="http://www.baidu.com"+ "//垂直滚动条\n//richTextBox1.ScrollBars = RichTextBoxScrollBars.Vertical;"
            + "//\n 水平滚动条";
            // 垂直滚动条
            richTextBox1.ScrollBars = RichTextBoxScrollBars.Vertical;
            // 水平滚动条
            //richTextBox1.WordWrap = false;  // 关闭自动换行 
            //richTextBox1.ScrollBars = RichTextBoxScrollBars.Horizontal;
            // both
            // richTextBox1.ScrollBars = RichTextBoxScrollBars.Both;

            // 设置文本样式
            richTextBox1.SelectionStart = 0;
            richTextBox1.SelectionLength = 5;
            richTextBox1.SelectionFont = new Font("楷体", 12, FontStyle.Bold);
            richTextBox1.SelectionColor = System.Drawing.Color.Green;

            // 段落样式
            richTextBox1.SelectionBullet = true;
            // 缩进
            richTextBox1.SelectionIndent = 20;
           

        }

        private void label1_Click(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            // 单击事件
            // 事件触发了标签的可见和不可见
            // MessageBox.Show("你单击了按钮，触发了事件");
            if (label1.Visible == true)
                label1.Visible = false;
            else
                label1.Visible = true;

        }


        private void button2_Click(object sender, EventArgs e)
        {
            // 按下回车键
            MessageBox.Show("按下回车键或单击按钮");

        }

        private void button3_Click(object sender, EventArgs e)
        {
            // 按下esc键
            MessageBox.Show("按下取消键");
        }

        private void textBox4_TextChanged(object sender, EventArgs e)
        {
            // 监听文本框的显示
            label1.Text = textBox4.Text;
        }

        private void richTextBox1_LinkClicked(object sender, LinkClickedEventArgs e)
        {
            // 网页跳转
            System.Diagnostics.Process.Start(e.LinkText);
        }
    }
}
