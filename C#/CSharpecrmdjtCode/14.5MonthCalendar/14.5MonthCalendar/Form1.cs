using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _14._5MonthCalendar
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void monthCalendar1_DateChanged(object sender, DateRangeEventArgs e)
        {
            // 开始
            label1.Text = monthCalendar1.SelectionStart.ToString();

            // 结束
            label2.Text = monthCalendar1.SelectionEnd.ToString();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            // 设置背景颜色
            monthCalendar1.TitleBackColor = System.Drawing.Color.Blue;

            monthCalendar1.TrailingForeColor = System.Drawing.Color.Red;

            monthCalendar1.TitleForeColor = System.Drawing.Color.Yellow;

            // 显示周数
            monthCalendar1.ShowWeekNumbers = true;

            // 显示多个月份
            monthCalendar1.CalendarDimensions = new Size(2, 2);

            // 设置粗体
            DateTime newdate = new DateTime(2019, 2, 19);
            monthCalendar1.AddBoldedDate(newdate);

            // 重绘
            monthCalendar1.UpdateBoldedDates();

            // 打印今日
            label1.Text = monthCalendar1.TodayDate.ToString();

            monthCalendar1.SelectionStart = monthCalendar1.TodayDate;

        }
    }
}
