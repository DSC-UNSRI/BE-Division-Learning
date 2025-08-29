export type Event = {
  id: number;
  location: string;
  start: string;
  cover: string;
};

export type EventResponse = {
  events: Event[]
};
export type EventCardProps = {
  id: number;
  title: string;
  date: string;
  description: string;
  image: string;
};