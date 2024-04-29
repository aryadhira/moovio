'use client'
import MovieList from "@/components/movielist";
import { useEffect } from "react";
import { useState } from "react";

const GetMovieList = (type) => {
  const [data,setData] = useState({})
  const [loading,setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/movies/getmovielist?list=${type}`)
      const result = await response.json()
      setData(result)
      setLoading(false)
    };
    fetchData();
  },[type])

  return {data,loading}
}

const Home = () => {
  const {data:latestdata,loading:latestloading} = GetMovieList("latest");
  const {data:imdbdata,loading:imdbloading} = GetMovieList("imdb");

  if(imdbloading || latestloading) {
      return (
          <div className="flex justify-center items-center h-screen">Loading...</div>
      );
  }

  return (
    <div>
      <MovieList moviedata={latestdata.data} header={"Latest Uploaded Movies"}/>
      <MovieList moviedata={imdbdata.data} header={"Top Rating Movies"}/>
    </div>
  );
}

export default Home;