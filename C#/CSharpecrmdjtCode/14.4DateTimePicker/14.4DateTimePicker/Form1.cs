using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _14._4DateTimePicker
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {

            dateTimePicker1.Format = DateTimePickerFormat.Long;
            // dateTimePicker1.CustomFormat = "MMMMM dd,yyyy-dddd";

            textBox1.Text = dateTimePicker1.Text;


            label1.Text = dateTimePicker1.Value.ToShortTimeString();
        }

        private void dateTimePicker1_ValueChanged(object sender, EventArgs e)
        {
            textBox1.Text = dateTimePicker1.Text;

        }
    }
}
