"use client"
import { useState, useEffect } from 'react'

const Page = ({ params }) => {
    const fallbackImage = 'https://placehold.co/600x900?text=Cover+Not+Available'

    const handleImgError = (e) => {
        e.target.src = fallbackImage; // Change the src to a fallback image
    };

    const [data,setData] = useState();
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/getmoviedetail?id=${params.id}`).
        then((res)=> res.json()).
        then((data) => {
            setData(data.data)
            setLoading(false)
        })
    },[])

    if(loading) {
        return (
            <div className="flex justify-center items-center h-screen">Loading...</div>
        );
    }

    const formatGenres = (genres) => {
        return genres.join(" - ");
    };

    return (
        <div className="w-full p-5 flex md:flex-row sm:flex-col">
            <div className="w-1/2 flex justify-center">
                <img width={600} height={600} src={data.cover} alt="..." onError={handleImgError}/>
            </div>
            <div className="p-10 w-1/2">
                <h1 className="font-bold text-2xl">{data.title}</h1>
                <h3 className="pt-5 text-red-500 text-md font-bold">Release Year : {data.year}</h3>
                <h3 className="pt-5 text-red-500 text-md font-bold">Genre : {formatGenres(data.category)}</h3>
                <h3 className="pt-5 pb-5 text-red-500">⭐️ {data.rating}</h3>
                <h2 className="font-bold text-xl">Synopsis:</h2>
                <p className="w-3/4 text-sm italic pt-5 pb-5">{data.synopsis}</p>
                <div className="flex gap-5">
                    <button className="rounded p-2 bg-red-500 w-20 shadow-xl hover:bg-red-400 hover:transition duration-300 ease-in-out hover:scale-125">720p</button>
                    <button className="rounded p-2 bg-red-500 w-20 shadow-xl hover:bg-red-400 hover:transition duration-300 ease-in-out hover:scale-125">1080p</button>
                </div>
            </div>
        </div>
    );
}

export default Page;