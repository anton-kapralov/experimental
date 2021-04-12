package kae.experimental.sumcalculator

import android.os.Bundle
import android.view.View
import android.widget.EditText
import androidx.appcompat.app.AppCompatActivity
import java.lang.Integer.parseInt

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }

    fun onCalculate(view: View) {
        val a = parseInt(findViewById<EditText>(R.id.a).text.toString())
        val b = parseInt(findViewById<EditText>(R.id.b).text.toString())
        val c = a + b
        findViewById<EditText>(R.id.c).setText(c.toString())
    }
}