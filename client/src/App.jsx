

import { useState } from "react"

function App() {

    const [url, setUrl] = useState("")
    const [shortCode, setShortCode] = useState("")

    const handleSubmit = async () => {
      const response = await fetch(
        "http://localhost:8000/shorten",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            url: url,
          }),
        }
      )

    const data = await response.json()

    setShortCode(data.shortCode)
    }

    const copyToClipboard = async () => {
      await navigator.clipboard.writeText(shortUrl)
    }

    const shortUrl = `http://localhost:8000/${shortCode}`


  return (
    <div className="relative min-h-screen bg-black flex items-center justify-center px-4 overflow-hidden">
<div className="absolute h-96 w-96 bg-blue-500/20 rounded-full blur-3xl -top-20 -left-20" />

<div className="absolute h-96 w-96 bg-purple-500/20 rounded-full blur-3xl bottom-0 right-0" />
      <div className="w-full max-w-2xl rounded-3xl border border-white/10 bg-white/5 backdrop-blur-xl p-10 shadow-2xl">
        <h1 className="text-6xl font-bold text-center text-white">
          URL Shortener
        </h1>

        <p className="text-slate-400 text-center mt-4 text-lg">
          Turn long links into clean, shareable URLs.
        </p>

        <div className="mt-10">
          <input
            type="text"
            placeholder="https://example.com"
            className="w-full rounded-2xl bg-black/30 border border-white/10 px-5 py-4 text-white outline-none focus:border-blue-500 transition"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              className="w-full rounded-2xl bg-black/30 border border-white/10 px-5 py-4 text-white outline-none focus:border-blue-500 transition"
          />

          <button onClick={handleSubmit} className="w-full mt-4 rounded-2xl bg-gradient-to-r from-blue-500 to-purple-500 text-white  font-semibold py-4 hover:scale-[1.02] transition-all duration-300"
          >
            Shorten URL
          </button>

          {shortCode && (
            <div className="mt-6 rounded-2xl border border-white/10 bg-black/30 p-4">
              <p className="text-slate-400 text-sm">
                Short URL
              </p>

              <p className="text-white mt-2">
                {shortUrl}
              </p>
            </div>
          )}
          <button
            onClick={copyToClipboard}
            className="mt-4 w-full rounded-xl bg-white text-black py-3 font-semibold hover:opacity-90 transition">
            Copy URL
          </button>
        </div>
      </div>
    </div>
  )
}

export default App