using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace _18iterate
{
    public partial class 泛型 : Form
    {
        public 泛型()
        {
            InitializeComponent();
        }

        private void 泛型_Load(object sender, EventArgs e)
        {

            // 实例化接口
            IGenericlnterface<System.ComponentModel.IListSource> factory =
                new Factory<System.Data.DataTable, System.ComponentModel.IListSource>();
            // 输出泛型的类型
            Console.WriteLine("类型:");
            Console.WriteLine(factory.CreateInstance().GetType().ToString());

            int i = Finder.Find<int>(new int[] { 1, 2, 3, 4, 5, 6, 7 }, 6);
            Console.WriteLine("6在数组中的位置" + i.ToString());


        }
    }

    public interface IGenericlnterface<T>
    {
        T CreateInstance();
    }
    // 实现泛型类
    public class Factory<T,TI>:IGenericlnterface<TI>where T:TI,new()
    {
        public TI CreateInstance()
        {
            return new T();
        }
    }

    public class Finder
    {
        public static int Find<T>(T[] items,T item)
        {
            for (int i = 0; i < items.Length; i++)
            {
                if (items[i].Equals(item))
                {
                    return i;  // 返回找见的索引
                }
            }
            return -1;
        }
    }

}
