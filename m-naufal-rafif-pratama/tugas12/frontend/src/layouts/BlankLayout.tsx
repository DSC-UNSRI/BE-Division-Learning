import type { ReactNode } from "react";

export default function BlankLayout({ children }: { children: ReactNode }) {
  return <div>{children}</div>;
}
