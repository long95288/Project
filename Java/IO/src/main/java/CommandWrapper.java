/**
 * @author longquanxiao
 * @date 2020/3/1
 */
import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;

public class CommandWrapper
{
    Process process;
    Thread in;
    Thread out;
    public CommandWrapper(Process process)
    {
        // 运行器
        this.process = process;
        // 输入流
        final InputStream inputStream = process.getInputStream();
        //
        final BufferedReader r=new BufferedReader(new InputStreamReader(inputStream));
        final byte[] buffer = new byte[1024];
        out = new Thread(){
            String line;
            int lineNumber=0;
            @Override
            public void run()
            {
                try {
                    while (true)
                    {
                        int count = inputStream.read(buffer);
                        System.out.println(lineNumber+":"+new String (buffer, 0, count-1));
                        line=r.readLine();
                        System.out.println(lineNumber+":"+line);
                        lineNumber++;
                    }
                }
                catch (Exception e)
                {

                }
            }
        };
        final BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        final OutputStream outputStream = process.getOutputStream();
        in = new Thread()
        {
            String line;
            @Override
            public void run()
            {
                try {
                    while (true)
                    {
                        outputStream.write((reader.readLine()+"/n").getBytes());
                        outputStream.flush();
                    }
                }
                catch (Exception e)
                {

                }
            }
        };
    }

    public void startIn()
    {
        in.start();
    }

    public void startOut()
    {
        out.start();
    }

    public void interruptIn()
    {
        in.interrupt();
    }

    public void interruptOut()
    {
        out.interrupt();
    }

    public static void main(String[] args)
    {
        try
        {
            CommandWrapper command = new CommandWrapper(Runtime.getRuntime().exec("native2ascii"));
            command.startIn();
            command.startOut();
        }
        catch (Exception e) {
            e.printStackTrace();
        }
    }

}
