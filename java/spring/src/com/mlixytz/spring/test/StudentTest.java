package com.mlixytz.spring.test;

import com.mlixytz.spring.beans.Student;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class StudentTest {

    public static void main(String[] args) {
        ApplicationContext applicationContext = new ClassPathXmlApplicationContext("beans-config.xml");
        Student student = (Student) applicationContext.getBean("stu1");
        System.out.println(student);
    }

}
