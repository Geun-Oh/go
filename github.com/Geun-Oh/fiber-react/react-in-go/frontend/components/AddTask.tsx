import React from "react";
import { mutate } from "swr";
import axios from "axios";

const fetcher = async (data: { title: string; body: string }) =>
  (await axios.post<{ title: string; body: string }>("/api/tasks", data)).data;

export default function AddTaskPage() {
  return (
    <div>
      {/* <input type="text" name="addTask" /> */}
      <button
        type="submit"
        name="addTask"
        onClick={async () => {
          const newTask = await fetcher({ title: "kjkj", body: "llll" });
          mutate("task", newTask);
        }}
      >
        클릭!
      </button>
    </div>
  );
}
