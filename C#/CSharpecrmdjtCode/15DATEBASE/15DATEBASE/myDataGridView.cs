using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Data.SqlClient;

namespace _15DATEBASE
{
    public partial class myDataGridView : Form
    {
        SqlConnection conn;
        SqlDataAdapter sda;
        DataSet ds = null;


        public myDataGridView()
        {
            // 调用属性文件初始化
            // 设计表中的文件函数
            InitializeComponent();
        }
        
        private void myDataGridView_Load(object sender, EventArgs e)
        {
            // 初始化数据库
            conn = new SqlConnection("server=.;database=EDUC;Integrated security=SSPI");
            conn.Open();
            if (conn.State == ConnectionState.Open)
            {
                // test1();
                test2();
            }
            else
            {
                MessageBox.Show("数据库连接失败");
            }
            test3();
            
        }

        private void test1()
        {
            // 实例化sqlDataAdapter对象
            sda = new SqlDataAdapter("select * from Student", conn);
            ds = new DataSet();
            sda.Fill(ds, "Student");
            dataGridView1.DataSource = ds.Tables[0];

        }
        private void test2()
        {
            sda = new SqlDataAdapter("select * from Class", conn);
            ds = new DataSet();
            sda.Fill(ds);
            dataGridView1.DataSource = ds.Tables[0];
            // 禁止显示行标题
            dataGridView1.RowHeadersVisible = false;
            conn.Close();

        }

        private void test3()
        {
            // 对表中数据进行操作
            // 4列
            dataGridView2.ColumnCount = 4;
            // 显示列标题
            dataGridView2.ColumnHeadersVisible = true;
            // 设置背景颜色
            DataGridViewCellStyle style = new DataGridViewCellStyle ();
            style.BackColor = Color.Chocolate;
            // 设置字体
            style.Font = new Font("Verdana", 10, FontStyle.Bold);
            // 应用样式
            dataGridView2.ColumnHeadersDefaultCellStyle = style;
            // 设置列名
            dataGridView2.Columns[0].Name = "编号";
            dataGridView2.Columns[1].Name = "姓名";
            dataGridView2.Columns[2].Name = "年龄";
            dataGridView2.Columns[3].Name = "性别";

            // 添加数据
            string[] row1 = new string[] {"001","小吕","28","男"};
            string[] row2 = new string[] { "002", "小张", "23", "女" };
            // 建立数组对象
            object[] rows = new object[] { row1, row2 };
            // 循环赋值
            foreach (string[] row in rows)
            {
                dataGridView2.Rows.Add(row);
            }


        }
        private void button1_Click(object sender, EventArgs e)
        {

            string msg = String.Format(
                "第{0}行,第{1}列",
                dataGridView1.CurrentCell.RowIndex,
                dataGridView1.CurrentCell.ColumnIndex);

            label1.Text = msg;
        }

        // 获得语句对应的数据表
        private DataTable dbconn(string strSql)
        {
            DataTable dtSelect = new DataTable();
            try
            {
                conn.Open();
                SqlDataAdapter sda2 = new SqlDataAdapter(strSql, conn);
               
                sda2.Fill(dtSelect);
               
                conn.Close();
                return dtSelect;
            }
            catch(Exception ex)
            {
                MessageBox.Show(ex.Message.ToString());
            }

            return dtSelect;
        }



        private Boolean dbUpdate()
        {
            string strSql = "select * from Class";
            DataTable dtUpdate = new DataTable();
            dtUpdate = this.dbconn(strSql);
            // 调用clear方法,清空数据
            dtUpdate.Rows.Clear();

            // 从表中获得数据
            DataTable dtShow = new DataTable();
            dtShow = (DataTable)this.dataGridView1.DataSource;
            //置入数据
            for (int i = 0; i < dtShow.Rows.Count; i++)
            {
                // 置入更新的数据
                dtUpdate.ImportRow(dtShow.Rows[i]);

            }
            try
            {
                // 这个builder有什么用？？
                SqlCommandBuilder cmdbuilder;
                cmdbuilder = new SqlCommandBuilder(this.sda);
                // 更新数据
                this.sda.Update(dtUpdate);

            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message.ToString());
                return false;
            }
            // 提交更改
            dtUpdate.AcceptChanges();
            return true;
        }

        private void button2_Click(object sender, EventArgs e)
        {
            if (dbUpdate())
            {
                MessageBox.Show("修改成功");
            }
            else
            {
                MessageBox.Show("修改失败");
            }
        }
    }
}
