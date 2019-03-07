using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _13._5
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
           panel1.Visible = false;
           richTextBox1.Text="姓名:233\n 性别:男";

            groupBox1.Text = "分组标题";

            // 设置图标


            // 设置选项卡为图标选项卡
            tabControl1.ImageList = imageList1;
            // 设置第二个选项卡是索引为0的图标
            tabPage2.ImageIndex = 0;
            tabPage2.Text = "选项卡2";
            Button tabbtn = new Button();
            tabbtn.Text = "选项卡按钮";
            tabPage2.Controls.Add(tabbtn);

            string Title = "新增选项卡" + (tabControl1.TabCount + 1).ToString();
            TabPage newPage = new TabPage(Title);
            tabControl1.TabPages.Add(newPage);


        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (textBox1.Text.Trim() =="233")
            {
                panel1.Show();
            }
            else
            {
                MessageBox.Show("没有233");
                textBox1.Text = "";
            }
        }

        private void richTextBox1_TextChanged(object sender, EventArgs e)
        {

        }

        private void button2_Click(object sender, EventArgs e)
        {

            string Title = "新增选项卡" + (tabControl1.TabCount + 1).ToString();
            TabPage newPage = new TabPage(Title);
            tabControl1.TabPages.Add(newPage);
        }

        private void button3_Click(object sender, EventArgs e)
        {
            if (tabControl1.SelectedIndex ==0)
            {
                MessageBox.Show("没有选择任何选项卡");
            }
            else
            {
                // 删除选项卡
                tabControl1.TabPages.Remove(tabControl1.SelectedTab);
            }
        }

        private void button4_Click(object sender, EventArgs e)
        {
            tabControl1.TabPages.Clear();
        }
    }
}
