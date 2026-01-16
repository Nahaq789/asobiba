use std::thread::JoinHandle;

fn main() {
    println!("start");
    let task_names: Vec<&str> = vec!["task1", "task2", "task3", "task4", "task5"];

    let mut handlers: Vec<JoinHandle<()>> = vec![];

    task_names.iter().for_each(|&n| {
        let task = std::thread::spawn(move || {
            println!("{:?} start", n);
            for _ in 0i64..5_000_000_000i64 {}
            println!("{:?} finish", n);
        });
        handlers.push(task);
    });

    for handle in handlers {
        handle.join().unwrap()
    }

    println!("all tasks finish")
}
