"use client"
import MovieList from "@/components/movielist";
import { useEffect, useState } from "react";

const BrowseMovie = () => {
    const [data, setData] = useState([])
    const [page, setPage] = useState(1)
    const [totalpage, setTotalPage] = useState(1)
    // const [loading, setLoading] = useState(true)

    const GetAllMovies = async () => {
        console.log("Fetching All Data...")
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/movies/getallmovies?page=${page}`)
        const result = await response.json()
        setData(result.data.movies)
        setTotalPage(result.data.totalmovies)
        // setLoading(false)
    }

    // if (loading) {
    //     return (
    //         <div className="flex justify-center items-center h-screen">
    //             <div className="loading"></div>
    //         </div>
    //     );
    // }

    useEffect(() => {
        console.log("test")
        GetAllMovies()
    }, [page])

    return (
        <div className="w-full flex flex-col justify-center items-center gap-5">
            <div className="flex flex-col pt-10 gap-3 w-full px-10 sm:px-32 md:px-52 lg:px-72 xl:px-96">
                <h2 className="text-white text-2xl font-bold">Search :</h2>
                <div className="flex flex-col md:flex-row gap-4">
                    <input type="text" className="w-full bg-slate-800 rounded-md p-5" />
                    <a href="" className="rounded p-3 md:w-52 bg-red-500 flex justify-center items-center font-bold text-md">Search</a>
                </div>
            </div>
            <div>
                <MovieList moviedata={data} header={"All Movies Data"} />
            </div>
            <div className="flex flex-row gap-5 p-5">
                <button className=" transition-all hover:text-red-500" onClick={() => setPage(page - 1)}>Previous</button>
                <h3>{page} of {totalpage}</h3>
                <button className=" transition-all hover:text-red-500" onClick={() => setPage(page + 1)}>Next</button>
            </div>
        </div>
    );
}

export default BrowseMovie;