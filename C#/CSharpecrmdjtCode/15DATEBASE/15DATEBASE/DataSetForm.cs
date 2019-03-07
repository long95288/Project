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
    public partial class DataSetForm : Form
    {
        public DataSetForm()
        {
            InitializeComponent();
        }

        /*
         *数据库的变量
         */
        SqlConnection conn;
        private void DataSetForm_Load(object sender, EventArgs e)
        {
            // 创建连接
            string connsql = "server=.;database=EDUC;integrated security=SSPI";
            conn = new SqlConnection(connsql);
            // 创建命令对象
            // SqlCommand cmd = new SqlCommand("select * from Class");

            conn.Open();
            if (conn.State==ConnectionState.Open)
            {

                // 测试合并功能
                //testMerge();
                // 测试复制功能
                testCopy();

            }
            else
            {
                MessageBox.Show("打开数据库失败！");
            }
        }


        private void testMerge()
        {
            // 创建Dataset
            DataSet ds = new DataSet();
            DataSet ds1 = new DataSet();
            // 适配器1
            SqlDataAdapter sda1 = new SqlDataAdapter("select * from Student", conn);
            //sda.SelectCommand = cmd;
            // 适配器2
            SqlDataAdapter sda2 = new SqlDataAdapter("select * from Class", conn);
            sda1.Fill(ds);
            // 使用数据填充
            sda2.Fill(ds1);
            // 合并数据
            ds1.Merge(ds, true, MissingSchemaAction.AddWithKey);

            // 将数据渲染
            dataGridView1.DataSource = ds1.Tables[0];
        }


        private void testCopy()
        {
            // 创建一个sqlcommand对象
            SqlCommand cmd = new SqlCommand("select * from Student", conn);
            // 创建一个适配器
            SqlDataAdapter sda = new SqlDataAdapter();
            sda.SelectCommand = cmd;
            // 创建一个数据集
            DataSet ds = new DataSet();
            // 填充数据
            sda.Fill(ds);
            // 渲染数据
            dataGridView1.DataSource = ds.Tables[0];

            // 使用复制功能
            DataSet ds2 = ds.Copy();
            dataGridView2.DataSource = ds2.Tables[0];


        }
    }
}
