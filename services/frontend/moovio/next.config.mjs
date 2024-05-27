/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: false,
    output: 'standalone',
    images: {
        remotePatterns: [
            {
                hostname: "yts.mx"
            }
        ]
    }
};

export default nextConfig;
