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
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        /*
         * 数据库连接的变量
         *
         */
        SqlConnection conn;
        SqlCommand cmd;
        private void button1_Click(object sender, EventArgs e)
        {
                try
                {
                    //string server = "server=.;database=test.mdf;uid=sa;pwd=";
                    string connsql = "server=.;database=EDUC;integrated security=SSPI";

                    conn = new SqlConnection(connsql);
                    conn.Open();
                    if (conn.State == ConnectionState.Open)
                    {
                        label1.Text = "连接成功";
                        button1.Enabled = false;
                        button2.Enabled = true;
                    // 调用开始函数
                    // start();
                    start2();
                    }
                    else
                    {
                        MessageBox.Show(conn.State.ToString());
                    }
                        

                }
                catch (Exception)
                {
                    MessageBox.Show("连接失败");

                }
        }

        private void button2_Click(object sender, EventArgs e)
        {
            try
            {
                // 断开连接
                conn.Dispose();

                if (conn.State == ConnectionState.Closed)
                {
                    label1.Text = "断开连接";
                    button2.Enabled = false;
                    button1.Enabled = true;
                }      
            }
            catch (Exception ex)
            {

                MessageBox.Show(ex.ToString());
            }
            
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        // 建立数据操作函数
        private void start()
        {
            // 建立查询对象
            cmd = new SqlCommand();
            // 查询对象和数据库连接起来
            cmd.Connection = conn;
            // 查询语句
            cmd.CommandText = "select * from Student";
            // 执行的语句类型，文本形式
            cmd.CommandType = CommandType.Text;

            // 存储过程类型
            // cmd.CommandType = CommandType.StoredProcedure;

            // 执行查询语句

            //int i = Convert.ToInt32(cmd.ExecuteNonQuery());
            //label2.Text = "执行结果"+i;
            // 读取执行结果
            SqlDataReader sdr = cmd.ExecuteReader();
            string tmp = "执行后的数据\n";
            while (sdr.Read())
            {
                // 打印执行结果
                //Console.WriteLine(sdr[0].ToString());
                for (int i = 0; i < sdr.FieldCount; i++)
                {
                    tmp += sdr[i].ToString();
                }
                tmp += "\n";
                Console.Write(tmp);
                //tmp += sdr[1].ToString();

            }
            richTextBox1.Text = tmp;


            // 使用ExecuteScaler()
            /*
            cmd.CommandText = "select count(*) from Student";
            cmd.CommandType = CommandType.Text;
            int num = Convert.ToInt32(cmd.ExecuteScalar());
            tmp += ("一共有"+num+"个学生");

            richTextBox1.Text = tmp;
            */
        }
        DataSet ds;
        SqlDataAdapter sda;
        private void start2()
        {
            // 创建命令对象
            SqlCommand cmd2 = new SqlCommand("select * from Student", conn);
            // 创建适配器
            sda = new SqlDataAdapter();
            // 设置语句
            sda.SelectCommand = cmd2;
            // 创建数据集
            ds = new DataSet();
            // 填充数据集
            sda.Fill(ds, "cs");
            // 设置数据集的表现
            dataGridView1.DataSource = ds.Tables[0];
        }

        private void button3_Click(object sender, EventArgs e)
        {
            DataTable dt = ds.Tables["cs"];
            sda.FillSchema(dt, SchemaType.Mapped);
            DataRow dr = dt.Rows.Find("5");
            // 设置这一行的值
            dr["sno"] = 20140101;
            dr["sname"] = "newline";
            dr["ssex"] = "男";
            dr["saddr"] = "NJLL";
            dr["sage"] = "33";
            dr["height"] = "1.54";
            // 实例化一个对象构建器
            SqlCommandBuilder cmdBuider = new SqlCommandBuilder(sda);
            // 调用更新函数
            sda.Update(dt);
        }

        private void dataGridView1_CellClick(object sender, DataGridViewCellEventArgs e)
        {
            label3.Text = dataGridView1.SelectedCells[0].Value.ToString();
            label4.Text = dataGridView1.SelectedCells[1].Value.ToString();
            label5.Text = dataGridView1.SelectedCells[2].Value.ToString();
            label6.Text = dataGridView1.SelectedCells[3].Value.ToString();
            label7.Text = dataGridView1.SelectedCells[4].Value.ToString();
            label8.Text = dataGridView1.SelectedCells[5].Value.ToString();


        }

        private void button4_Click(object sender, EventArgs e)
        {
            DataSetForm dsf = new DataSetForm();
            dsf.Show();
            this.Hide();
        }

        private void button5_Click(object sender, EventArgs e)
        {
            myDataGridView my = new myDataGridView();
            my.Show();
            this.Hide();
        }
    }
}
