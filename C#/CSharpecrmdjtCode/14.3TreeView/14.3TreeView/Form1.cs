using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _14._3TreeView
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            // 建立3个父节点
            TreeNode tn1 = treeView1.Nodes.Add("名称");
            TreeNode tn2 = treeView1.Nodes.Add("性别");
            TreeNode tn3 = treeView1.Nodes.Add("年龄");

            // 建立子节点
            TreeNode Ntn1 = new TreeNode("1");
            TreeNode Ntn2 = new TreeNode("2");
            TreeNode Ntn3 = new TreeNode("3");

            //添加入父节点中
            tn1.Nodes.Add(Ntn1);
            tn1.Nodes.Add(Ntn2);
            tn1.Nodes.Add(Ntn3);

            // sexnode
            TreeNode sex_nan = new TreeNode("男");
            TreeNode sex_nv = new TreeNode("女");

            tn2.Nodes.Add(sex_nan);
            tn2.Nodes.Add(sex_nv);

            treeView1.ImageIndex = 2;
            treeView1.SelectedImageIndex = 1;


        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (treeView1.SelectedNode.Text=="名称")
            {

            }
            else
            {
                treeView1.Nodes.Remove(treeView1.SelectedNode);
            }
        }

        private void treeView1_AfterCheck(object sender, TreeViewEventArgs e)
        {
            label1.Text = "当前选中节点" + e.Node.Text;
        }

        private void button2_Click(object sender, EventArgs e)
        {
            if (textBox1.Text =="")
            {
                MessageBox.Show("请输入节点名称");

            }
            else
            {
                if (true)
                {
                    // 添加节点
                    TreeNode tmp = new TreeNode(textBox1.Text);
                    treeView1.SelectedNode.Nodes.Add(tmp);
                    textBox1.Text = "";
                }
            }
        }
    }
}
