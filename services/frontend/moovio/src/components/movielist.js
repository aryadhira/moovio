"use client"
import { useRouter } from "next/navigation";

const MovieList = ({moviedata, header}) => {

    const fallbackImage = 'https://placehold.co/600x900?text=Cover+Not+Available'

    const handleImgError = (e) => {
        e.target.src = fallbackImage; // Change the src to a fallback image
    };
    const router = useRouter()
    const gotodetail = (id) =>{
        router.push(`/detail/${id}`)
    };

    return (
        <div>
            <h1 className="text-white font-bold text-2xl px-10 pt-10">{header}</h1>
            <div className="grid xl:grid-cols-8 lg:grid-cols-5 md:grid-cols-4 sm:grid-cols-3 grid-cols-2 gap-6 p-10">
            {moviedata.data.map(movie => {
                return (
                    <div key={movie.id} onClick={() => gotodetail(`${movie.id}`)} className="p-3 flex flex-col rounded-md shadow-lg bg-red-500 transition-all hover:scale-105">
                        <img src={movie.cover} width={300} height={300} alt={movie.title} onError={handleImgError}/>
                        <h3 className="font-bold text-md pt-4">{movie.title}</h3>
                        <h5 className="font-bold text-white">({movie.year})</h5>
                        <h5 className="font-bold text-white">⭐️ {movie.rating}</h5>
                    </div>
                )
            })}
        </div>
        </div>
        
        
    );
        
}

export default MovieList;