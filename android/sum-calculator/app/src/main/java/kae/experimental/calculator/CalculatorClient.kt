package kae.experimental.calculator

import android.util.Log
import androidx.lifecycle.MutableLiveData
import io.ktor.client.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.http.content.*
import kae.experimental.calculator.CalculatorServiceProtos.SumRequest
import kae.experimental.calculator.CalculatorServiceProtos.SumResponse

private const val SERVICE_URL =
    "http://10.0.0.182:8080/twirp/kae.experimental.calculator.CalculatorService"

class CalculatorClient() {

    suspend fun sum(addends: Iterable<Int>, result: MutableLiveData<Int>) {
        val request = SumRequest.newBuilder().addAllAddends(addends).build()

        try {
            val httpResponse =
                HttpClient().use {
                    it.post<HttpResponse>("$SERVICE_URL/Sum") {
                        body = ByteArrayContent(
                            request.toByteArray(),
                            ContentType.Application.ProtoBuf
                        )
                    }
                }
            val response =
                SumResponse.parseFrom(httpResponse.readBytes())

            result.postValue(response.sum)
        } catch (e: Exception) {
            Log.e(null, null, e)
        }
    }

}