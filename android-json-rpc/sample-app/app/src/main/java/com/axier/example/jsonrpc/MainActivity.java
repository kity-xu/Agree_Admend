package com.axier.example.jsonrpc;

import android.os.AsyncTask;
import android.os.Bundle;
import android.support.v7.app.ActionBarActivity;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import com.axier.jsonrpclibrary.JSONRPCClient;
import com.axier.jsonrpclibrary.JSONRPCException;
import com.axier.jsonrpclibrary.JSONRPCParams;

import org.json.JSONObject;


public class MainActivity extends ActionBarActivity {

    public String URL = "";
    public String EXAMPLE_SUCCESS_CALL = "https://raw.githubusercontent.com/axierjhtjz/android-json-rpc/master/success.json";
    public String EXAMPLE_ERROR_CALL = "https://raw.githubusercontent.com/axierjhtjz/android-json-rpc/master/error.json";
    public String EXAMPLE_ads_CALL = "http://192.168.13.175:8888/rpc";
//    public String EXAMPLE_ads_CALL = "http://10.30.0.12:8888/rpc";
//    public String EXAMPLE_ads_CALL = "http://mpro.sinaapp.com/my/jzdw.php?hex=0&lac=10328&cid=26997&map=0";

    /**
     * Just a example of how to use it against a WS. In this case we are just fetching data from the
     * url above.
     */
    public String EXAMPLE_METHOD_NAME = "login";
    public String METHOD_NAME = "Pr2Protocol.PrintEx";
    public String METHOD_NAME1 = "";
//    public String EXAMPLE_PARAM_1 = "user";
//    public String EXAMPLE_PARAM_2 = "password";
//    public String PARAM = "message";

    public TextView mResponseArea;
    public TextView mResponseArea2;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        mResponseArea = (TextView) findViewById(R.id.response_area);
        mResponseArea2 = (TextView) findViewById(R.id.response_area2);

        Button successBtn = (Button) findViewById(R.id.success_btn);
        Button errorBtn = (Button) findViewById(R.id.error_btn);
        Button adsBtn = (Button) findViewById(R.id.ads_btn);

        successBtn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                URL = EXAMPLE_SUCCESS_CALL;
                METHOD_NAME1 = EXAMPLE_METHOD_NAME;
                new MakeJSONRpcCallTask().execute();
            }
        });

        errorBtn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                URL = EXAMPLE_ERROR_CALL;
                METHOD_NAME1 = EXAMPLE_METHOD_NAME;
                new MakeJSONRpcCallTask().execute();
            }
        });
        adsBtn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                URL = EXAMPLE_ads_CALL;
                METHOD_NAME1 = METHOD_NAME;
                new MakeJSONRpcCallTask().execute();
            }
        });

        new Thread(new Runnable() {
            @Override
            public void run() {

            }
        }).start();
    }

    public class MakeJSONRpcCallTask extends AsyncTask<Void, Void, JSONObject> {
        @Override
        protected JSONObject doInBackground(Void... params) {
            JSONRPCClient client = JSONRPCClient.create(URL, JSONRPCParams.Versions.VERSION_2);
            client.setConnectionTimeout(20000);
            client.setSoTimeout(20000);
            try {
                JSONObject jsonObj = new JSONObject();
//                jsonObj.put(EXAMPLE_PARAM_1, "myuser");
//                jsonObj.put(EXAMPLE_PARAM_2, "mypassword");
//                jsonObj.put("Timeout", 10000);
//                jsonObj.put("Con", "abcdefg");
//                jsonObj.put("LpicAppData", "");
                return client.callJSONObject(METHOD_NAME1, jsonObj);
            } catch (JSONRPCException rpcEx) {
                rpcEx.printStackTrace();
//            } catch (JSONException jsEx) {
//                jsEx.printStackTrace();
            }
            return null;
        }

        @Override
        protected void onPostExecute(JSONObject result) {
            super.onPostExecute(result);
            if (result != null && mResponseArea != null) {
                mResponseArea.setText(result.toString());
                mResponseArea2.setText(result.toString());
            }
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        return super.onOptionsItemSelected(item);
    }
}