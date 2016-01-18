package cn.com.agree.ab.amend;

import android.content.Intent;
import android.content.SharedPreferences;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.TextView;

public class SettingActivity extends AppCompatActivity {
    TextView btn;
    EditText ip_edt, port_edt, servlet_edt;
    private SharedPreferences sharedPrefrences;
    private SharedPreferences.Editor editor;
    private static final String FILENAME = "filename";
    public static final int MODE_APPEND = 0x8000;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_setting);


        ip_edt = (EditText)findViewById(R.id.id_ip);
        port_edt = (EditText)findViewById(R.id.id_port);
        servlet_edt = (EditText)findViewById(R.id.id_servlet);

        sharedPrefrences = this.getSharedPreferences(FILENAME, MODE_APPEND);
        setSharedPreferences();


        btn = (TextView)findViewById(R.id.id_set_btn);
        btn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                saveSharedPreferences();
                Intent intent = new Intent(SettingActivity.this, ManagerActivity.class);
                startActivity(intent);
                finish();
            }
        });

    }

    void setSharedPreferences(){
        String r_server = sharedPrefrences.getString("server", "");
        String r_ip = sharedPrefrences.getString("ip", "");
        String r_port= sharedPrefrences.getString("port", "");
        servlet_edt.setText(r_server);
        ip_edt.setText(r_ip);
        port_edt.setText(r_port);

    }
    void saveSharedPreferences(){
        editor = this.getSharedPreferences(FILENAME, MODE_APPEND).edit();
        String server=servlet_edt.getText().toString();
        String ip=ip_edt.getText().toString();
        String port=port_edt.getText().toString();
        editor.putString("server", server);
        editor.putString("ip", ip);
        editor.putString("port", port);
        editor.commit();
    }
}
