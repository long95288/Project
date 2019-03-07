using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _13._6
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void toolStripComboBox1_Click(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            this.toolStripProgressBar2.Minimum = 0;
            this.toolStripProgressBar2.Maximum = 5000;
            this.toolStripProgressBar2.Step = 2;
            for (int i = 0; i <= 4999; i++)
            {
                // 按照属性增加
                this.toolStripProgressBar2.PerformStep();
            }
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            this.toolStripStatusLabel1.Text = DateTime.Now.ToShortDateString();

        }
    }
}
