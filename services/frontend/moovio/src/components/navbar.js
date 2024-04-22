import Link from "next/link";
const Navbar = () => {
    return (
        <div className="w-full flex justify-between items-center bg-red-500 p-5">
            <Link href="/">
                <h1 className="text-white font-bold text-2xl">Moovio.</h1>
            </Link>
            <Link href="/browse">
                <h1 className="text-white font-bold text-md">Browse Movie</h1>
            </Link>
        </div>
    );
}

export default Navbar;