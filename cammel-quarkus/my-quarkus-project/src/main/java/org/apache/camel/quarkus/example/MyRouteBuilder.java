import org.apache.camel.builder.RouteBuilder;

public class MyRouteBuilder extends RouteBuilder {
    @Override
    public void configure() throws Exception {
        from("platform-http:/hello").setBody(constant("Hello From Camel Quarkus !"));
    }
}
