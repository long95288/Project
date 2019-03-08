using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;

namespace _20File
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (textBox1.Text == string.Empty)
            {
                MessageBox.Show("文件名不能为空");

            }
            else
            {
                if (File.Exists(textBox1.Text))
                {
                    MessageBox.Show("文件已经存在了");
                }
                else
                {
                    // 如果不存在则创建文件
                    File.Create(textBox1.Text);
                    MessageBox.Show("文件" + textBox1.Text + "创建成功");

                }
            }
        }
    }
}
