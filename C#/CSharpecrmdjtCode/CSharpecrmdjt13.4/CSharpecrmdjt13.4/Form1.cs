using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace CSharpecrmdjt13._4
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            comboBox1.DropDownStyle = ComboBoxStyle.DropDownList;
            comboBox1.Items.Add("select 1");
            comboBox1.Items.Add("select 2");
            comboBox1.Items.Add("select 3");

            comboBox2.DropDownStyle = ComboBoxStyle.DropDown;
            comboBox2.Items.Add("select 1");
            comboBox2.Items.Add("select 2");
            comboBox2.Items.Add("select 3");

            comboBox3.DropDownStyle = ComboBoxStyle.Simple;
            comboBox3.Items.Add("select 1");
            comboBox3.Items.Add("select 2");
            comboBox3.Items.Add("select 3");

            // 数字选择框
            numericUpDown1.Maximum = 100;
            numericUpDown1.Minimum = 1;
            // 显示小数点后两位
            numericUpDown1.DecimalPlaces = 2;

            listBox1.HorizontalScrollbar = true;
            listBox1.ScrollAlwaysVisible = true;
            // 选择模式
            listBox1.SelectionMode = SelectionMode.MultiExtended;


        }

        private void button1_Click(object sender, EventArgs e)
        {
            comboBox2.SelectAll();
        }

        private void comboBox1_SelectedValueChanged(object sender, EventArgs e)
        {
            // 选项发生变化时监听的动作
            label1.Text = comboBox1.Text;
        }

        private void checkBox1_CheckedChanged(object sender, EventArgs e)
        {

        }

        private void checkBox1_CheckStateChanged(object sender, EventArgs e)
        {
            label1.Text = "状态发生变化";
            if (checkBox1.CheckState == CheckState.Checked)
                label1.Text = "选中";
            else
                label1.Text = "没被选中";
       
        }

        private void checkBox1_Click(object sender, EventArgs e)
        {
            if (checkBox1.CheckState == CheckState.Checked)
                MessageBox.Show("被选中");
            else
                MessageBox.Show("没被选中");
        }

        private void radioButton1_CheckedChanged(object sender, EventArgs e)
        {

        }

        private void radioButton1_Click(object sender, EventArgs e)
        {
            if (radioButton1.Checked == true)
                label1.Text = "单选框被选中";
            else
                label1.Text = "单选框没被选中";

            radioButton2.Checked = false;
        }

        private void radioButton2_Click(object sender, EventArgs e)
        {
            radioButton1.Checked = false;

        }

        private void listBox1_SelectedIndexChanged(object sender, EventArgs e)
        {

        }

        private void addbtn_Click(object sender, EventArgs e)
        {
            if (textBox1.Text == "")
                MessageBox.Show("输入数据");
            else
            {
                listBox1.Items.Add(textBox1.Text);
                textBox1.Text = "";
            }
        }

        private void delbtn_Click(object sender, EventArgs e)
        {
            if (listBox1.SelectedItems.Count ==0)
            {
                MessageBox.Show("选择数据后删除");
            }
            else
            {
                // 删除所有的被选项
                while(listBox1.SelectedItems.Count!=0)
                    listBox1.Items.Remove(listBox1.SelectedItem);
            }
        }
    }
}
