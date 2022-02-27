import core.cache as cache
import core.client as client
import core.util as util


def fetch_resources():
    resource = (
        client.Client()
        .set_headers({"Content-Type": "application/json"})
        .get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
    )

    currency = cache.get_currency_cache_data()
    result = None

    try:
        currency["data"]["IDR_USD"]
    except KeyError as e:
        result = (
            client.Client()
            .set_headers({"Content-Type": "application/json"})
            .get(
                "https://free.currconv.com/api/v7/convert?q=IDR_USD&apiKey=6c686cb7c5ac91832327"
            )
        )

        cache.put_cache_file(
            util.get_currency_cache_file(), "IDR_USD", result["results"]["IDR_USD"]
        )

    if result != None:
        currency["data"]["IDR_USD"] = result["results"]["IDR_USD"]

    result = []

    for val in resource:
        result.append(
            {
                "uuid": val["uuid"],
                "komoditas": val["komoditas"],
                "area_provinsi": val["area_provinsi"],
                "area_kota": val["area_kota"],
                "size": val["size"],
                "price": val["price"],
                "price_usd": None
                if val["price"] == None
                else int(val["price"]) * currency["data"]["IDR_USD"]["val"],
                "tgl_parsed": val["tgl_parsed"],
                "timestamp": val["timestamp"],
            }
        )

    return result
