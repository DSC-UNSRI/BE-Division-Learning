export default function Footer() {
  return (
    <footer className="bg-gray-800 text-gray-200 py-4 mt-8">
      <div className="max-w-7xl mx-auto text-center text-sm">
        © {new Date().getFullYear()} BackendLearn12. All rights reserved.
      </div>
    </footer>
  );
}
