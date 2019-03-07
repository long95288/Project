using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _14._1
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            //项目根目录
            string rootPath = Application.StartupPath.Substring(0, Application.StartupPath.Substring(0,Application.
                StartupPath.LastIndexOf("\\")).LastIndexOf("\\"));

            // 第一张图片的地址
            string Path1 =rootPath + @"\02.png";
            // 第二张图片地址
            string Path2 = rootPath + @"\03.png";
            // 创建
            Image img1 = Image.FromFile(Path1, true);
            Image img2 = Image.FromFile(Path2, true);

            imageList1.Images.Add(img1);
            imageList1.Images.Add(img2);

            // 设置图片大小
            imageList1.ImageSize = new Size(200, 165);
            pictureBox1.Width = 200;
            pictureBox1.Height = 165;

        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (imageList1.Images.Count != 0)
            {
                // 显示第一个图像
                pictureBox1.Image = imageList1.Images[0];

            }
            
        }

        private void button2_Click(object sender, EventArgs e)
        {
            if (imageList1.Images.Count != 0)
            {
                // 显示第二张图像
                pictureBox1.Image = imageList1.Images[1];
            }
           
        }

        private void button3_Click(object sender, EventArgs e)
        {
            imageList1.Images.RemoveAt(0);
        }
    }
}
