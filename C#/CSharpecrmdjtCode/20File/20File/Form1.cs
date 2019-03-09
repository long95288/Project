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

        private void button2_Click(object sender, EventArgs e)
        {
            if (textBox2.Text == string.Empty)
            {
                MessageBox.Show("文件夹名称不能为空");
            }
            else
            {
                if (Directory.Exists(textBox2.Text))
                {

                    MessageBox.Show("文件夹已经存在");
                }
                else
                {
                    Directory.CreateDirectory(textBox2.Text);
                }
            }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            if (textBox3.Text == string.Empty)
            {
                MessageBox.Show("文件夹名称不能为空");
            }
            else
            {
                FileInfo finfo = new FileInfo(textBox3.Text);
                if (finfo.Exists)
                {
                    MessageBox.Show("该文件已经存在");

                }
                else
                {
                    finfo.Create();
                    // finfo.CopyTo("D:\\test.html", true);


                }
            }
        }

        private void button4_Click(object sender, EventArgs e)
        {
            if (textBox4.Text == string.Empty)
            {
                MessageBox.Show("文件夹名称不能为空");

            }
            else
            {
                DirectoryInfo dinfo = new DirectoryInfo(textBox4.Text);
                if (dinfo.Exists)
                {
                    MessageBox.Show("该文件夹已经存在");

                }
                else
                {
                    dinfo.Create();
                }
            }
            
        }

        private void button5_Click(object sender, EventArgs e)
        {
            File.Copy("E:\test.txt", "D:\test.html");
        }

        private void button6_Click(object sender, EventArgs e)
        {
            File.Move("E:\test.txt", "F:\test.html");
        }

        private void button7_Click(object sender, EventArgs e)
        {
            File.Delete("F:\test.html");
            File.Delete("D:\test.html");

        }

        private void button8_Click(object sender, EventArgs e)
        {
            if (openFileDialog1.ShowDialog() == DialogResult.OK)
            {
                textBox5.Text = openFileDialog1.FileName;
                FileInfo finfo = new FileInfo(textBox5.Text);
                string strCTime, strLATime, strLWTime, strName, strFName, strDName, strISRead;
                long lgLength;
                strCTime = finfo.CreationTime.ToShortDateString();
                strLATime = finfo.LastAccessTime.ToShortDateString();
                strLWTime = finfo.LastWriteTime.ToShortDateString();
                strName = finfo.Name;
                strFName = finfo.FullName;
                strDName = finfo.DirectoryName;
                strISRead = finfo.IsReadOnly.ToString();
                lgLength = finfo.Length;
                string info = "文件信息:\n 创建时间:" + strCTime + "\n上次访问时间:" +
                    strLATime + "\n上次写入时间" + strLWTime + "\n文件名称:" + strName +
                    "\n完整目录" + strFName + "\n 完整路径" + strDName + "\n是否只读:" +
                    strISRead + "\n文件长度:" + lgLength;
                MessageBox.Show(info);
            }
        }
    }
}
