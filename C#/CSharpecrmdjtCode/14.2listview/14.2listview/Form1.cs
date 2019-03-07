using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _14._2listview
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (textBox1.Text=="")
            {
                MessageBox.Show("项目不能为空");
            }
            else
            {
                listView1.Items.Add(textBox1.Text.Trim());
                textBox1.Text = "";
            }
        }

        private void button2_Click(object sender, EventArgs e)
        {
           
            if (listView1.SelectedItems.Count!= 0)
            {
                listView1.Items.RemoveAt(listView1.SelectedItems[0].Index);
                listView1.SelectedItems.Clear();
            }
            else
            {
                MessageBox.Show("没有选择一个选项");
            }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            if (listView1.Items.Count==0)
            {
                MessageBox.Show("没有项目");
            }
            else
            {
                // 删除所有
                listView1.Items.Clear();
            }
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            listView1.Items[2].Selected = true;
            listView1.Items[0].ImageIndex = 0;
            listView1.Items[1].ImageIndex = 1;
            listView1.Items[2].ImageIndex = 2;

            // 新建项目分组
            listView1.Groups.Add(new ListViewGroup("测试", HorizontalAlignment.Left));

            // 分配
            listView1.Items[0].Group = listView1.Groups[0];
        }
    }
}
